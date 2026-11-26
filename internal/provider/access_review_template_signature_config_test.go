package provider

import (
	"context"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"testing"

	tfTypes "github.com/conductorone/terraform-provider-conductorone/internal/provider/types"
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/models/operations"
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	tfresource "github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// These tests lock IGA-3923: removing signature_config from a
// conductorone_access_review_template must clear it on the API instead of
// silently keeping the prior value.

func accessReviewTemplateSchema(t *testing.T) schema.Schema {
	t.Helper()
	resp := &resource.SchemaResponse{}
	NewAccessReviewTemplateResource().Schema(context.Background(), resource.SchemaRequest{}, resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("schema diagnostics: %v", resp.Diagnostics)
	}
	return resp.Schema
}

// TestAccessReviewTemplateSignatureConfigIsNotComputed guards the overlay entry.
// A Computed parent makes Terraform keep prior state when the block is removed.
func TestAccessReviewTemplateSignatureConfigIsNotComputed(t *testing.T) {
	attr, ok := accessReviewTemplateSchema(t).Attributes["signature_config"].(schema.SingleNestedAttribute)
	if !ok {
		t.Fatalf("signature_config is not a SingleNestedAttribute")
	}
	if !attr.IsOptional() {
		t.Errorf("signature_config must be Optional")
	}
	if attr.IsComputed() {
		t.Errorf("signature_config must not be Computed; removing the block would keep prior state")
	}
	for name, leaf := range attr.Attributes {
		if !leaf.IsOptional() || !leaf.IsComputed() {
			t.Errorf("signature_config.%s must stay Optional+Computed so partially configured blocks keep server values", name)
		}
	}
}

func signatureConfigStateObject(t *testing.T, s schema.Schema, sig attr.Value) tfsdk.State {
	t.Helper()
	ctx := context.Background()
	state := tfsdk.State{Schema: s, Raw: tftypes.NewValue(s.Type().TerraformType(ctx), nil)}
	if diags := state.SetAttribute(ctx, path.Root("signature_config"), sig); diags.HasError() {
		t.Fatalf("set state: %v", diags)
	}
	return state
}

func TestAddSignatureConfigRemovalToUpdateMask(t *testing.T) {
	s := accessReviewTemplateSchema(t)
	sigType, ok := s.Attributes["signature_config"].GetType().(types.ObjectType)
	if !ok {
		t.Fatalf("signature_config type is %T, want types.ObjectType", s.Attributes["signature_config"].GetType())
	}
	setSig := types.ObjectValueMust(sigType.AttrTypes, map[string]attr.Value{
		"meaning_of_signature": types.StringValue("I approve"),
		"require_signature":    types.BoolValue(true),
		"step_up_provider_id":  types.StringValue(""),
		"tsp_url":              types.StringValue(""),
	})
	nullSig := types.ObjectNull(sigType.AttrTypes)

	const baseMask = "description,displayName"
	cases := []struct {
		name      string
		prior     attr.Value
		planned   *tfTypes.ReviewSignatureConfig
		wantMask  string
		wantNoReq bool
	}{
		{name: "removed", prior: setSig, planned: nil, wantMask: baseMask + ",signatureConfig"},
		{name: "never set", prior: nullSig, planned: nil, wantMask: baseMask},
		{name: "still set", prior: setSig, planned: &tfTypes.ReviewSignatureConfig{}, wantMask: baseMask},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			mask := baseMask
			req := &operations.C1APIAccessreviewV1AccessReviewTemplateServiceUpdateRequest{
				AccessReviewTemplateServiceUpdateRequest: &shared.AccessReviewTemplateServiceUpdateRequest{UpdateMask: &mask},
			}
			data := &AccessReviewTemplateResourceModel{SignatureConfig: tc.planned}
			diags := addSignatureConfigRemovalToUpdateMask(context.Background(), signatureConfigStateObject(t, s, tc.prior), data, req)
			if diags.HasError() {
				t.Fatalf("diagnostics: %v", diags)
			}
			if got := *req.AccessReviewTemplateServiceUpdateRequest.UpdateMask; got != tc.wantMask {
				t.Errorf("mask = %q, want %q", got, tc.wantMask)
			}
		})
	}
}

// fakeTemplateAPI models the access review template endpoints with the field
// mask semantics of the ConductorOne API: a masked path absent from the body is
// cleared, and naming signatureConfig without the e-signature feature is a 403.
type fakeTemplateAPI struct {
	mu         sync.Mutex
	esignature bool
	templates  map[string]map[string]any
	masks      []string
	nextID     int
}

func (f *fakeTemplateAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f.mu.Lock()
	defer f.mu.Unlock()

	const prefix = "/api/v1/access_review_template"
	id := strings.TrimPrefix(strings.TrimPrefix(r.URL.Path, prefix), "/")
	var body map[string]any
	if r.Body != nil {
		_ = json.NewDecoder(r.Body).Decode(&body)
	}
	switch {
	case !strings.HasPrefix(r.URL.Path, prefix):
		http.NotFound(w, r)
	case r.Method == http.MethodPost && id == "":
		if _, ok := body["signatureConfig"]; ok && !f.esignature {
			writeJSON(w, http.StatusForbidden, map[string]any{"code": 7, "message": "e-signature feature is not enabled"})
			return
		}
		f.nextID++
		newID := fmt.Sprintf("tmpl-%d", f.nextID)
		body["id"] = newID
		f.templates[newID] = body
		writeJSON(w, http.StatusOK, map[string]any{"accessReviewTemplate": body})
	case r.Method == http.MethodGet:
		tmpl, ok := f.templates[id]
		if !ok {
			writeJSON(w, http.StatusNotFound, map[string]any{"code": 5, "message": "not found"})
			return
		}
		writeJSON(w, http.StatusOK, map[string]any{"accessReviewTemplate": tmpl})
	case r.Method == http.MethodPost:
		tmpl, ok := f.templates[id]
		if !ok {
			writeJSON(w, http.StatusNotFound, map[string]any{"code": 5, "message": "not found"})
			return
		}
		mask, _ := body["updateMask"].(string)
		f.masks = append(f.masks, mask)
		in, _ := body["accessReviewTemplate"].(map[string]any)
		for _, p := range strings.Split(mask, ",") {
			if p == "signatureConfig" && !f.esignature {
				writeJSON(w, http.StatusForbidden, map[string]any{"code": 7, "message": "e-signature feature is not enabled"})
				return
			}
		}
		for _, p := range strings.Split(mask, ",") {
			if v, ok := in[p]; ok && v != nil {
				tmpl[p] = v
			} else {
				delete(tmpl, p)
			}
		}
		writeJSON(w, http.StatusOK, map[string]any{"accessReviewTemplate": tmpl})
	case r.Method == http.MethodDelete:
		delete(f.templates, id)
		writeJSON(w, http.StatusOK, map[string]any{})
	default:
		http.Error(w, "unsupported", http.StatusMethodNotAllowed)
	}
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func (f *fakeTemplateAPI) only(t *testing.T) map[string]any {
	t.Helper()
	f.mu.Lock()
	defer f.mu.Unlock()
	if len(f.templates) != 1 {
		t.Fatalf("expected one template on the fake API, got %d", len(f.templates))
	}
	for _, v := range f.templates {
		return v
	}
	return nil
}

func (f *fakeTemplateAPI) checkSignatureConfig(t *testing.T, want bool) tfresource.TestCheckFunc {
	return func(*terraform.State) error {
		_, got := f.only(t)["signatureConfig"]
		if got != want {
			return fmt.Errorf("API signatureConfig present = %v, want %v", got, want)
		}
		return nil
	}
}

func (f *fakeTemplateAPI) checkLastMaskHasSignatureConfig(want bool) tfresource.TestCheckFunc {
	return func(*terraform.State) error {
		f.mu.Lock()
		defer f.mu.Unlock()
		if len(f.masks) == 0 {
			return fmt.Errorf("no update was sent")
		}
		last := f.masks[len(f.masks)-1]
		if got := strings.Contains(","+last+",", ",signatureConfig,"); got != want {
			return fmt.Errorf("last update mask %q names signatureConfig = %v, want %v", last, got, want)
		}
		return nil
	}
}

const signatureConfigLifecycleChildEnv = "C1_TEST_SIGNATURE_CONFIG_CHILD"

var (
	fakeTemplateAPIOnce sync.Once
	sharedFakeAPI       = &fakeTemplateAPI{}
	fakeTemplateAPIErr  error
)

// startFakeTemplateAPI returns the process-wide fake, reset for a new test.
// The server is shared because Go caches trusted roots for the process after the
// first TLS handshake, so a second server certificate would not be trusted.
// The provider forces https; the certificate directory is exported so terraform
// and released provider binaries started as subprocesses trust it as well.
func startFakeTemplateAPI(t *testing.T, esignature bool) *fakeTemplateAPI {
	t.Helper()
	fakeTemplateAPIOnce.Do(func() {
		srv := httptest.NewTLSServer(sharedFakeAPI)
		dir, err := os.MkdirTemp("", "c1-fake-api-")
		if err != nil {
			fakeTemplateAPIErr = err
			return
		}
		certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: srv.Certificate().Raw})
		if err := os.WriteFile(filepath.Join(dir, "fake-api.pem"), certPEM, 0o600); err != nil {
			fakeTemplateAPIErr = err
			return
		}
		certDirs := os.Getenv("SSL_CERT_DIR")
		if certDirs == "" {
			certDirs = "/etc/ssl/certs"
		}
		for k, v := range map[string]string{
			"SSL_CERT_DIR":               dir + string(os.PathListSeparator) + certDirs,
			"CONDUCTORONE_ACCESS_TOKEN":  "fake-token",
			"CONDUCTORONE_SERVER_URL":    srv.URL,
			"CONDUCTORONE_TENANT_DOMAIN": "",
			"CONDUCTORONE_OIDC_TOKEN":    "",
		} {
			if err := os.Setenv(k, v); err != nil {
				fakeTemplateAPIErr = err
				return
			}
		}
	})
	if fakeTemplateAPIErr != nil {
		t.Fatal(fakeTemplateAPIErr)
	}
	sharedFakeAPI.mu.Lock()
	defer sharedFakeAPI.mu.Unlock()
	sharedFakeAPI.esignature = esignature
	sharedFakeAPI.templates = map[string]map[string]any{}
	sharedFakeAPI.masks = nil
	return sharedFakeAPI
}

// runSignatureConfigLifecycleInChild re-runs the named test in a fresh process
// so the fake API's certificate is trusted regardless of earlier TLS use.
func runSignatureConfigLifecycleInChild(t *testing.T) bool {
	t.Helper()
	if os.Getenv(signatureConfigLifecycleChildEnv) == t.Name() {
		return false
	}
	if os.Getenv("TF_ACC") == "" {
		t.Skip("set TF_ACC=1 to run Terraform against the fake API")
	}
	cmd := exec.Command(os.Args[0], "-test.run=^"+t.Name()+"$", "-test.v", "-test.timeout=10m")
	cmd.Env = append(os.Environ(), signatureConfigLifecycleChildEnv+"="+t.Name())
	out, err := cmd.CombinedOutput()
	t.Log(string(out))
	if err != nil {
		t.Fatalf("child test failed: %v", err)
	}
	return true
}

func templateConfig(sigAttr, sigBody, description string) string {
	sig := ""
	if sigBody != "" {
		sig = fmt.Sprintf("  %s = %s\n", sigAttr, sigBody)
	}
	return fmt.Sprintf(`
resource "conductorone_access_review_template" "test" {
  display_name = "IGA-3923"
  description  = %q
%s}
`, description, sig)
}

const signatureBlock = `{
    require_signature    = true
    meaning_of_signature = "I attest"
  }`

func noDriftAfterApply(t *testing.T) tfresource.ConfigPlanChecks {
	return tfresource.ConfigPlanChecks{PostApplyPostRefresh: []plancheck.PlanCheck{ExpectEmptyPlanWithDriftLog(t)}}
}

func expectAction(action plancheck.ResourceActionType) tfresource.ConfigPlanChecks {
	return tfresource.ConfigPlanChecks{PreApply: []plancheck.PlanCheck{
		plancheck.ExpectResourceAction("conductorone_access_review_template.test", action),
	}}
}

func withPostApply(t *testing.T, c tfresource.ConfigPlanChecks) tfresource.ConfigPlanChecks {
	c.PostApplyPostRefresh = []plancheck.PlanCheck{ExpectEmptyPlanWithDriftLog(t)}
	return c
}

// TestAccessReviewTemplateSignatureConfigLifecycle drives Terraform through the
// four shapes against the fake API: set at create, removed, added later, and
// removed with an explicit null.
func TestAccessReviewTemplateSignatureConfigLifecycle(t *testing.T) {
	if runSignatureConfigLifecycleInChild(t) {
		return
	}
	fake := startFakeTemplateAPI(t, true)
	tfresource.Test(t, tfresource.TestCase{
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []tfresource.TestStep{
			{
				Config:           templateConfig("signature_config", signatureBlock, "a"),
				ConfigPlanChecks: noDriftAfterApply(t),
				Check:            fake.checkSignatureConfig(t, true),
			},
			{
				Config:           templateConfig("signature_config", "", "a"),
				ConfigPlanChecks: withPostApply(t, expectAction(plancheck.ResourceActionUpdate)),
				Check: tfresource.ComposeTestCheckFunc(
					fake.checkSignatureConfig(t, false),
					fake.checkLastMaskHasSignatureConfig(true),
					tfresource.TestCheckNoResourceAttr("conductorone_access_review_template.test", "signature_config.%"),
				),
			},
			{
				Config:           templateConfig("signature_config", signatureBlock, "a"),
				ConfigPlanChecks: withPostApply(t, expectAction(plancheck.ResourceActionUpdate)),
				Check: tfresource.ComposeTestCheckFunc(
					fake.checkSignatureConfig(t, true),
					tfresource.TestCheckResourceAttr("conductorone_access_review_template.test", "signature_config.require_signature", "true"),
				),
			},
			{
				Config:           templateConfig("signature_config", "null", "a"),
				ConfigPlanChecks: withPostApply(t, expectAction(plancheck.ResourceActionUpdate)),
				Check:            fake.checkSignatureConfig(t, false),
			},
		},
	})
}

// TestAccessReviewTemplateSignatureConfigNeverSet covers a template that never
// configures signature_config on a tenant without the e-signature feature:
// updates must not name signatureConfig in the mask (IGA-2595).
func TestAccessReviewTemplateSignatureConfigNeverSet(t *testing.T) {
	if runSignatureConfigLifecycleInChild(t) {
		return
	}
	fake := startFakeTemplateAPI(t, false)
	tfresource.Test(t, tfresource.TestCase{
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []tfresource.TestStep{
			{
				Config:           templateConfig("signature_config", "", "a"),
				ConfigPlanChecks: noDriftAfterApply(t),
				Check:            fake.checkSignatureConfig(t, false),
			},
			{
				Config:           templateConfig("signature_config", "", "b"),
				ConfigPlanChecks: withPostApply(t, expectAction(plancheck.ResourceActionUpdate)),
				Check: tfresource.ComposeTestCheckFunc(
					fake.checkSignatureConfig(t, false),
					fake.checkLastMaskHasSignatureConfig(false),
				),
			},
		},
	})
}

func upgradeSteps(t *testing.T, fake *fakeTemplateAPI, version, oldAttr string, setAtCreate bool) []tfresource.TestStep {
	oldBody, newBody := "", ""
	if setAtCreate {
		oldBody, newBody = signatureBlock, signatureBlock
	}
	external := map[string]tfresource.ExternalProvider{
		"conductorone": {Source: "conductorone/conductorone", VersionConstraint: version},
	}
	steps := []tfresource.TestStep{
		{
			ExternalProviders: external,
			Config:            templateConfig(oldAttr, oldBody, "a"),
			Check:             fake.checkSignatureConfig(t, setAtCreate),
		},
		{
			ProtoV6ProviderFactories: testAccProviderFactories,
			Config:                   templateConfig("signature_config", newBody, "a"),
			ConfigPlanChecks: tfresource.ConfigPlanChecks{PreApply: []plancheck.PlanCheck{
				ExpectEmptyPlanWithDriftLog(t),
			}},
			Check: fake.checkSignatureConfig(t, setAtCreate),
		},
	}
	if setAtCreate {
		steps = append(steps, tfresource.TestStep{
			ProtoV6ProviderFactories: testAccProviderFactories,
			Config:                   templateConfig("signature_config", "", "a"),
			ConfigPlanChecks:         withPostApply(t, expectAction(plancheck.ResourceActionUpdate)),
			Check: tfresource.ComposeTestCheckFunc(
				fake.checkSignatureConfig(t, false),
				fake.checkLastMaskHasSignatureConfig(true),
			),
		})
	} else {
		steps = append(steps, tfresource.TestStep{
			ProtoV6ProviderFactories: testAccProviderFactories,
			Config:                   templateConfig("signature_config", "", "b"),
			ConfigPlanChecks:         withPostApply(t, expectAction(plancheck.ResourceActionUpdate)),
			Check:                    fake.checkLastMaskHasSignatureConfig(false),
		})
	}
	return steps
}

// TestAccessReviewTemplateSignatureConfigUpgrade applies with a released
// provider, then switches to this build: existing state must plan empty, and
// removal must clear the API value afterwards. v1.4.3 named the attribute
// review_signature_config; it became signature_config in v1.5.0.
func TestAccessReviewTemplateSignatureConfigUpgrade(t *testing.T) {
	if runSignatureConfigLifecycleInChild(t) {
		return
	}
	for _, tc := range []struct {
		version, oldAttr string
	}{
		{version: "1.4.3", oldAttr: "review_signature_config"},
		{version: "1.8.3", oldAttr: "signature_config"},
	} {
		for _, set := range []bool{true, false} {
			t.Run(fmt.Sprintf("v%s/set=%v", tc.version, set), func(t *testing.T) {
				fake := startFakeTemplateAPI(t, set)
				tfresource.Test(t, tfresource.TestCase{Steps: upgradeSteps(t, fake, tc.version, tc.oldAttr, set)})
			})
		}
	}
}
