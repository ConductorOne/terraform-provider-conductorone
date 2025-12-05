resource "conductorone_integration_artifactory" "artifactory" {
  app_id = conductorone_app.artifactory.id
  user_ids = [
    conductorone_user.admin.id
  ]
  artifactory_base_url                  = "..."
  artifactory_reference_token           = "..."
  log_level                             = "..."
  otel_collector_endpoint               = "..."
  otel_collector_endpoint_tls_cert      = "..."
  otel_collector_endpoint_tls_cert_path = "..."
  otel_collector_endpoint_tls_insecure  = false
  otel_logging_disabled                 = false
  otel_tracing_disabled                 = false
}
