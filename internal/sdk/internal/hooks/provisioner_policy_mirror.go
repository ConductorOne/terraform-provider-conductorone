package hooks

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

// provisionerPolicyMirror copies appEntitlement.provisionerPolicy into the
// synthetic read-only provisionerPolicyDeprecated field so the generated
// Custom App Entitlement resource can expose the deprecated provisioner_policy
// attribute as a mirror of provision_policy.
type provisionerPolicyMirror struct{}

var mirrorOps = map[string]bool{
	"c1.api.app.v1.AppEntitlements.Create": true,
	"c1.api.app.v1.AppEntitlements.Get":    true,
	"c1.api.app.v1.AppEntitlements.Update": true,
}

func (provisionerPolicyMirror) AfterSuccess(hookCtx AfterSuccessContext, res *http.Response) (*http.Response, error) {
	if !mirrorOps[hookCtx.OperationID] || res == nil || res.Body == nil {
		return res, nil
	}
	raw, err := io.ReadAll(res.Body)
	_ = res.Body.Close()
	if err != nil {
		return nil, err
	}
	out := raw
	var doc map[string]any
	if json.Unmarshal(raw, &doc) == nil {
		if view, ok := doc["appEntitlementView"].(map[string]any); ok {
			if ent, ok := view["appEntitlement"].(map[string]any); ok {
				if p, ok := ent["provisionerPolicy"]; ok {
					ent["provisionerPolicyDeprecated"] = p
					if b, err := json.Marshal(doc); err == nil {
						out = b
					}
				}
			}
		}
	}
	res.Body = io.NopCloser(bytes.NewReader(out))
	res.ContentLength = int64(len(out))
	return res, nil
}
