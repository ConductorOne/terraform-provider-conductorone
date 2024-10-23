resource "conductorone_integration_jd_edwards" "jd_edwards" {
  app_id = conductorone_app.jd_edwards.id
  user_ids = [
    conductorone_user.admin.id
  ]
  jdedwards_ais_url  = "..."
  jdedwards_username = "..."
  jdedwards_password = "..."
  jdedwards_env      = "..."
}
