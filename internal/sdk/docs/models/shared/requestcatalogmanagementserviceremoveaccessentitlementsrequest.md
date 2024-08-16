# RequestCatalogManagementServiceRemoveAccessEntitlementsRequest

The RequestCatalogManagementServiceRemoveAccessEntitlementsRequest message is used to remove access entitlements from a request catalog.
 The access entitlements are used to determine which users can view the request catalog.


## Fields

| Field                                                           | Type                                                            | Required                                                        | Description                                                     |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| `AccessEntitlements`                                            | [][AppEntitlementRef](../../models/shared/appentitlementref.md) | :heavy_minus_sign:                                              | The list of access entitlements to remove from the catalog.     |