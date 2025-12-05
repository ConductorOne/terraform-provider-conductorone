resource "conductorone_integration_tenable_vm" "tenable_vm" {
  app_id = conductorone_app.tenable_vm.id
  user_ids = [
    conductorone_user.admin.id
  ]
  tenable_vm_access_key = "..."
  tenable_vm_secret_key = "..."
}
