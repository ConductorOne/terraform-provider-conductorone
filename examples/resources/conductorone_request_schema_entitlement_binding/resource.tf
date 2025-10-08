resource "conductorone_request_schema_entitlement_binding" "my_request_schema_entitlement_binding" {
  entitlement_refs = [
    {
      app_id = "...my_app_id..."
      id     = "...my_id..."
    }
  ]
  request_schema_id = "...my_request_schema_id..."
}