resource "terraform_connector_credential" "my_connectorcredential" {
  app_id                                      = "...my_app_id..."
  connector_id                                = "...my_connector_id..."
  connector_service_rotate_credential_request = {}
}