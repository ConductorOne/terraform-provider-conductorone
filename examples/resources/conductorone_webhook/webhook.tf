resource "conductorone_webhook" "new_webhook" {
  display_name = "New TF sourced webhook"
  url          = "https://<webhook_url>"
  description  = "New webhook created using Terraform"
}
