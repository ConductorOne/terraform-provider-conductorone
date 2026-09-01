package provider

import (
	"context"

	tfTypes "github.com/conductorone/terraform-provider-conductorone/internal/provider/types"
	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

type appEntitlementKey struct {
	appID string
	id    string
}

func appEntitlementKeyFromTerraform(ref tfTypes.AppEntitlementRef) (appEntitlementKey, bool) {
	if ref.AppID.IsNull() || ref.AppID.IsUnknown() || ref.ID.IsNull() || ref.ID.IsUnknown() {
		return appEntitlementKey{}, false
	}
	return appEntitlementKey{appID: ref.AppID.ValueString(), id: ref.ID.ValueString()}, true
}

func appEntitlementKeyFromShared(ref shared.AppEntitlementRef) (appEntitlementKey, bool) {
	if ref.AppID == nil || ref.ID == nil {
		return appEntitlementKey{}, false
	}
	return appEntitlementKey{appID: *ref.AppID, id: *ref.ID}, true
}

func appEntitlementRefDifference(before, after []tfTypes.AppEntitlementRef) (added, removed []tfTypes.AppEntitlementRef) {
	beforeKeys := make(map[appEntitlementKey]struct{}, len(before))
	for _, ref := range before {
		if key, ok := appEntitlementKeyFromTerraform(ref); ok {
			beforeKeys[key] = struct{}{}
		}
	}
	afterKeys := make(map[appEntitlementKey]struct{}, len(after))
	for _, ref := range after {
		if key, ok := appEntitlementKeyFromTerraform(ref); ok {
			afterKeys[key] = struct{}{}
		}
	}
	for _, ref := range after {
		if key, ok := appEntitlementKeyFromTerraform(ref); ok {
			if _, exists := beforeKeys[key]; !exists {
				added = append(added, ref)
			}
		}
	}
	for _, ref := range before {
		if key, ok := appEntitlementKeyFromTerraform(ref); ok {
			if _, exists := afterKeys[key]; !exists {
				removed = append(removed, ref)
			}
		}
	}
	return added, removed
}

func managedAppEntitlementRefs(managed []tfTypes.AppEntitlementRef, live []shared.AppEntitlementRef) []tfTypes.AppEntitlementRef {
	liveKeys := make(map[appEntitlementKey]struct{}, len(live))
	for _, ref := range live {
		if key, ok := appEntitlementKeyFromShared(ref); ok {
			liveKeys[key] = struct{}{}
		}
	}
	refs := make([]tfTypes.AppEntitlementRef, 0, len(managed))
	for _, ref := range managed {
		if key, ok := appEntitlementKeyFromTerraform(ref); ok {
			if _, exists := liveKeys[key]; exists {
				refs = append(refs, ref)
			}
		}
	}
	return refs
}

func (r *AccessProfileRequestableEntriesResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	var prior resource.SchemaResponse
	r.Schema(ctx, resource.SchemaRequest{}, &prior)
	prior.Schema.Version = 0

	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema: &prior.Schema,
			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var state *AccessProfileRequestableEntriesResourceModel
				resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
				if resp.Diagnostics.HasError() {
					return
				}

				state.AppEntitlements = nil
				resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
			},
		},
	}
}
