package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// Hand-maintained; listed in .genignore.
//
// conductorone_custom_app_entitlement used to expose the policy as a computed
// `provisioner_policy` (the API field name). It now exposes a writable
// `provision_policy`. This wrapper keeps the old attribute as a deprecated
// read-only mirror of `provision_policy` so existing configurations and state
// keep working, without editing the generated resource.
//
// The generated resource decodes plan and state into a struct that has no
// `provisioner_policy` field, so each call hands it a copy of the request with
// that attribute removed and adds it back to the resulting state.

const (
	provisionPolicyAttr     = "provision_policy"
	provisionerPolicyAttr   = "provisioner_policy"
	deprovisionerPolicyAttr = "deprovisioner_policy"

	provisionerPolicyDeprecation = "Use provision_policy instead. provisioner_policy is a read-only copy of provision_policy."
)

type customAppEntitlementInner interface {
	resource.Resource
	resource.ResourceWithConfigure
	resource.ResourceWithImportState
}

type customAppEntitlementAliasResource struct {
	inner customAppEntitlementInner
}

var (
	_ resource.Resource                = &customAppEntitlementAliasResource{}
	_ resource.ResourceWithConfigure   = &customAppEntitlementAliasResource{}
	_ resource.ResourceWithImportState = &customAppEntitlementAliasResource{}
)

// NewCustomAppEntitlementAliasResource wraps the generated custom app
// entitlement resource with the deprecated provisioner_policy attribute.
func NewCustomAppEntitlementAliasResource() resource.Resource {
	return &customAppEntitlementAliasResource{inner: NewCustomAppEntitlementResource().(customAppEntitlementInner)}
}

func (r *customAppEntitlementAliasResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	r.inner.Metadata(ctx, req, resp)
}

func (r *customAppEntitlementAliasResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.inner.Configure(ctx, req, resp)
}

func (r *customAppEntitlementAliasResource) innerSchema(ctx context.Context) (schema.Schema, diag.Diagnostics) {
	resp := &resource.SchemaResponse{}
	r.inner.Schema(ctx, resource.SchemaRequest{}, resp)
	return resp.Schema, resp.Diagnostics
}

func (r *customAppEntitlementAliasResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema, resp.Diagnostics = r.innerSchema(ctx)
	if resp.Diagnostics.HasError() {
		return
	}
	policy, ok := resp.Schema.Attributes[provisionPolicyAttr].(schema.SingleNestedAttribute)
	if !ok {
		resp.Diagnostics.AddError("Unexpected schema", provisionPolicyAttr+" is not a single nested attribute")
		return
	}

	mirror := computedOnlyNested(policy)
	mirror.Description = "Deprecated copy of `provision_policy`."
	mirror.MarkdownDescription = "Deprecated copy of `provision_policy`."
	mirror.DeprecationMessage = provisionerPolicyDeprecation

	attrs := make(map[string]schema.Attribute, len(resp.Schema.Attributes)+1)
	for name, attr := range resp.Schema.Attributes {
		attrs[name] = attr
	}
	attrs[provisionerPolicyAttr] = mirror
	resp.Schema.Attributes = attrs
}

func (r *customAppEntitlementAliasResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	inner, diags := r.innerSchema(ctx)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	plan, d1 := toInner(ctx, inner, req.Plan.Raw)
	config, d2 := toInner(ctx, inner, req.Config.Raw)
	resp.Diagnostics.Append(append(d1, d2...)...)
	if resp.Diagnostics.HasError() {
		return
	}

	innerReq := resource.CreateRequest{
		Config:       tfsdk.Config{Schema: inner, Raw: config},
		Plan:         tfsdk.Plan{Schema: inner, Raw: plan},
		Identity:     req.Identity,
		ProviderMeta: req.ProviderMeta,
	}
	innerResp := &resource.CreateResponse{
		State:    tfsdk.State{Schema: inner, Raw: tftypes.NewValue(inner.Type().TerraformType(ctx), nil)},
		Identity: resp.Identity,
		Private:  resp.Private,
	}
	r.inner.Create(ctx, innerReq, innerResp)

	resp.Diagnostics.Append(innerResp.Diagnostics...)
	resp.Identity = innerResp.Identity
	resp.Private = innerResp.Private
	resp.Diagnostics.Append(setOuterState(ctx, &resp.State, innerResp.State.Raw)...)
}

func (r *customAppEntitlementAliasResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	inner, diags := r.innerSchema(ctx)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	state, d := toInner(ctx, inner, req.State.Raw)
	resp.Diagnostics.Append(d...)
	if resp.Diagnostics.HasError() {
		return
	}

	innerReq := resource.ReadRequest{
		State:              tfsdk.State{Schema: inner, Raw: state},
		Identity:           req.Identity,
		Private:            req.Private,
		ProviderMeta:       req.ProviderMeta,
		ClientCapabilities: req.ClientCapabilities,
	}
	innerResp := &resource.ReadResponse{
		State:    tfsdk.State{Schema: inner, Raw: state},
		Identity: resp.Identity,
		Private:  resp.Private,
	}
	r.inner.Read(ctx, innerReq, innerResp)

	resp.Diagnostics.Append(innerResp.Diagnostics...)
	resp.Identity = innerResp.Identity
	resp.Private = innerResp.Private
	resp.Deferred = innerResp.Deferred
	resp.Diagnostics.Append(setOuterState(ctx, &resp.State, innerResp.State.Raw)...)
}

func (r *customAppEntitlementAliasResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	inner, diags := r.innerSchema(ctx)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	plan, d1 := toInner(ctx, inner, req.Plan.Raw)
	config, d2 := toInner(ctx, inner, req.Config.Raw)
	state, d3 := toInner(ctx, inner, req.State.Raw)
	resp.Diagnostics.Append(append(append(d1, d2...), d3...)...)
	if resp.Diagnostics.HasError() {
		return
	}
	for _, policy := range []string{provisionPolicyAttr, deprovisionerPolicyAttr} {
		var err error
		if state, err = dropStaleArms(config, state, policy); err != nil {
			resp.Diagnostics.AddError("Unexpected value", fmt.Sprintf("preparing %s for update: %s", policy, err))
			return
		}
	}

	innerReq := resource.UpdateRequest{
		Config:       tfsdk.Config{Schema: inner, Raw: config},
		Plan:         tfsdk.Plan{Schema: inner, Raw: plan},
		State:        tfsdk.State{Schema: inner, Raw: state},
		Identity:     req.Identity,
		ProviderMeta: req.ProviderMeta,
		Private:      req.Private,
	}
	// The framework starts an update response from the plan.
	innerResp := &resource.UpdateResponse{
		State:    tfsdk.State{Schema: inner, Raw: plan},
		Identity: resp.Identity,
		Private:  resp.Private,
	}
	r.inner.Update(ctx, innerReq, innerResp)

	resp.Diagnostics.Append(innerResp.Diagnostics...)
	resp.Identity = innerResp.Identity
	resp.Private = innerResp.Private
	resp.Diagnostics.Append(setOuterState(ctx, &resp.State, innerResp.State.Raw)...)
}

func (r *customAppEntitlementAliasResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	inner, diags := r.innerSchema(ctx)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	state, d := toInner(ctx, inner, req.State.Raw)
	resp.Diagnostics.Append(d...)
	if resp.Diagnostics.HasError() {
		return
	}

	innerReq := resource.DeleteRequest{
		State:        tfsdk.State{Schema: inner, Raw: state},
		Identity:     req.Identity,
		ProviderMeta: req.ProviderMeta,
		Private:      req.Private,
	}
	innerResp := &resource.DeleteResponse{
		State:    tfsdk.State{Schema: inner, Raw: state},
		Identity: resp.Identity,
		Private:  resp.Private,
	}
	r.inner.Delete(ctx, innerReq, innerResp)

	resp.Diagnostics.Append(innerResp.Diagnostics...)
	resp.Identity = innerResp.Identity
	resp.Private = innerResp.Private
}

func (r *customAppEntitlementAliasResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	inner, diags := r.innerSchema(ctx)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	innerResp := &resource.ImportStateResponse{
		State:    tfsdk.State{Schema: inner, Raw: tftypes.NewValue(inner.Type().TerraformType(ctx), nil)},
		Identity: resp.Identity,
	}
	r.inner.ImportState(ctx, req, innerResp)

	resp.Diagnostics.Append(innerResp.Diagnostics...)
	resp.Identity = innerResp.Identity
	resp.Diagnostics.Append(setOuterState(ctx, &resp.State, innerResp.State.Raw)...)
}

// toInner removes provisioner_policy from an outer-schema value so it matches
// the generated resource's schema. Null and unknown values pass through typed
// for the inner schema.
func toInner(ctx context.Context, inner schema.Schema, raw tftypes.Value) (tftypes.Value, diag.Diagnostics) {
	var diags diag.Diagnostics
	innerType := inner.Type().TerraformType(ctx)
	if raw.IsNull() || !raw.IsKnown() {
		return tftypes.NewValue(innerType, nil), diags
	}
	attrs, err := objectAttrs(raw)
	if err != nil {
		diags.AddError("Unexpected value", fmt.Sprintf("reading resource value: %s", err))
		return raw, diags
	}
	delete(attrs, provisionerPolicyAttr)
	return tftypes.NewValue(innerType, attrs), diags
}

// setOuterState stores an inner-schema value in state, with provisioner_policy
// copied from provision_policy. A null value (resource removed) stays null.
func setOuterState(ctx context.Context, state *tfsdk.State, innerRaw tftypes.Value) diag.Diagnostics {
	var diags diag.Diagnostics
	outerType := state.Schema.Type().TerraformType(ctx)
	if innerRaw.IsNull() {
		state.Raw = tftypes.NewValue(outerType, nil)
		return diags
	}
	attrs, err := objectAttrs(innerRaw)
	if err != nil {
		diags.AddError("Unexpected value", fmt.Sprintf("reading resource state: %s", err))
		return diags
	}
	attrs[provisionerPolicyAttr] = attrs[provisionPolicyAttr]
	state.Raw = tftypes.NewValue(outerType, attrs)
	return diags
}

// dropStaleArms nulls, in the prior state, every arm of the named policy that
// configuration does not set, when configuration sets exactly one arm.
//
// The API models a policy as a oneof and rejects a body with two arms. The
// generated Update seeds its model from prior state and overlays the plan, but an
// arm left unset in configuration is unknown in the plan and does not overwrite
// the state, so switching arms would send the old arm next to the new one. A
// configuration with no arm, or more than one, is left alone.
func dropStaleArms(config, state tftypes.Value, policy string) (tftypes.Value, error) {
	cfgAttrs, err := objectAttrs(config)
	if err != nil {
		return state, err
	}
	stateAttrs, err := objectAttrs(state)
	if err != nil {
		return state, err
	}
	cfgPolicy, statePolicy := cfgAttrs[policy], stateAttrs[policy]
	if cfgPolicy.IsNull() || !cfgPolicy.IsKnown() || statePolicy.IsNull() || !statePolicy.IsKnown() {
		return state, nil
	}
	cfgArms, err := objectAttrs(cfgPolicy)
	if err != nil {
		return state, err
	}
	stateArms, err := objectAttrs(statePolicy)
	if err != nil {
		return state, err
	}

	var selected string
	n := 0
	for name, arm := range cfgArms {
		if !arm.IsNull() && arm.IsKnown() {
			selected = name
			n++
		}
	}
	if n != 1 {
		return state, nil
	}
	for name, arm := range stateArms {
		if name != selected {
			stateArms[name] = tftypes.NewValue(arm.Type(), nil)
		}
	}
	stateAttrs[policy] = tftypes.NewValue(statePolicy.Type(), stateArms)
	return tftypes.NewValue(state.Type(), stateAttrs), nil
}

// objectAttrs returns a copy of an object value's attributes. Value.As hands
// back the value's own map, so editing it in place would corrupt the original.
func objectAttrs(v tftypes.Value) (map[string]tftypes.Value, error) {
	var src map[string]tftypes.Value
	if err := v.As(&src); err != nil {
		return nil, err
	}
	attrs := make(map[string]tftypes.Value, len(src))
	for name, val := range src {
		attrs[name] = val
	}
	return attrs, nil
}

// computedOnlyNested returns a copy of a nested attribute that is read-only:
// computed, with no validators, plan modifiers or defaults at any depth.
func computedOnlyNested(a schema.SingleNestedAttribute) schema.SingleNestedAttribute {
	attrs := make(map[string]schema.Attribute, len(a.Attributes))
	for name, attr := range a.Attributes {
		attrs[name] = computedOnly(attr)
	}
	return schema.SingleNestedAttribute{
		Computed:            true,
		Attributes:          attrs,
		CustomType:          a.CustomType,
		Description:         a.Description,
		MarkdownDescription: a.MarkdownDescription,
	}
}

// computedOnly converts the attribute kinds the policy schema uses. Any other
// kind panics so that a regenerated schema cannot silently drop a flag.
func computedOnly(attr schema.Attribute) schema.Attribute {
	switch a := attr.(type) {
	case schema.SingleNestedAttribute:
		return computedOnlyNested(a)
	case schema.StringAttribute:
		return schema.StringAttribute{Computed: true, CustomType: a.CustomType, Description: a.Description, MarkdownDescription: a.MarkdownDescription}
	case schema.BoolAttribute:
		return schema.BoolAttribute{Computed: true, CustomType: a.CustomType, Description: a.Description, MarkdownDescription: a.MarkdownDescription}
	case schema.ListAttribute:
		return schema.ListAttribute{Computed: true, ElementType: a.ElementType, CustomType: a.CustomType, Description: a.Description, MarkdownDescription: a.MarkdownDescription}
	case schema.MapAttribute:
		return schema.MapAttribute{Computed: true, ElementType: a.ElementType, CustomType: a.CustomType, Description: a.Description, MarkdownDescription: a.MarkdownDescription}
	}
	panic(fmt.Sprintf("computedOnly: unsupported attribute type %T", attr))
}
