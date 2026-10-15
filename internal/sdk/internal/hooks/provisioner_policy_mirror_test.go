package hooks

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"
)

func mirrorResponse(t *testing.T, operationID, body string) string {
	t.Helper()
	res := &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(strings.NewReader(body)), ContentLength: int64(len(body))}
	out, err := provisionerPolicyMirror{}.AfterSuccess(AfterSuccessContext{HookContext{OperationID: operationID}}, res)
	if err != nil {
		t.Fatal(err)
	}
	data, err := io.ReadAll(out.Body)
	if err != nil {
		t.Fatal(err)
	}
	if out.ContentLength != int64(len(data)) {
		t.Fatalf("ContentLength = %d, body is %d bytes", out.ContentLength, len(data))
	}
	return string(data)
}

func TestProvisionerPolicyMirror(t *testing.T) {
	const withPolicy = `{"appEntitlementView":{"appEntitlement":{"id":"e1","provisionerPolicy":{"delegated":{"appId":"a2","entitlementId":"e2"}}}}}`

	for _, op := range []string{"c1.api.app.v1.AppEntitlements.Create", "c1.api.app.v1.AppEntitlements.Get", "c1.api.app.v1.AppEntitlements.Update"} {
		t.Run(op, func(t *testing.T) {
			var doc struct {
				View struct {
					Entitlement map[string]json.RawMessage `json:"appEntitlement"`
				} `json:"appEntitlementView"`
			}
			if err := json.Unmarshal([]byte(mirrorResponse(t, op, withPolicy)), &doc); err != nil {
				t.Fatal(err)
			}
			got, want := doc.View.Entitlement["provisionerPolicyDeprecated"], doc.View.Entitlement["provisionerPolicy"]
			if len(want) == 0 || string(got) != string(want) {
				t.Fatalf("provisionerPolicyDeprecated = %s, want %s", got, want)
			}
		})
	}

	t.Run("other operations are untouched", func(t *testing.T) {
		if got := mirrorResponse(t, "c1.api.app.v1.AppEntitlements.List", withPolicy); got != withPolicy {
			t.Fatalf("body = %s", got)
		}
	})
	t.Run("a response without a policy is untouched", func(t *testing.T) {
		const without = `{"appEntitlementView":{"appEntitlement":{"id":"e1"}}}`
		if got := mirrorResponse(t, "c1.api.app.v1.AppEntitlements.Get", without); got != without {
			t.Fatalf("body = %s", got)
		}
	})
	t.Run("a body that is not JSON is passed through", func(t *testing.T) {
		if got := mirrorResponse(t, "c1.api.app.v1.AppEntitlements.Get", "not json"); got != "not json" {
			t.Fatalf("body = %s", got)
		}
	})
}
