resource "conductorone_app_owner" "test_okta_app_admin_myself" {
  app_id = data.conductorone_app.test_okta.id
  user_ids = [
    data.conductorone_user.my_user.id
  ]
}
