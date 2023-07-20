resource "conductorone_integration_sentry" "sentry" {
  app_id          = conductorone_app.sentry.id
  sentry_org_slug = "..."
  sentry_token    = "..."
}
