resource "conductorone_integration_lucidchart" "lucidchart" {
  app_id = conductorone_app.lucidchart.id
  user_ids = [
    conductorone_user.admin.id
  ]
  lucidchart_authorization_token = "..."
}
