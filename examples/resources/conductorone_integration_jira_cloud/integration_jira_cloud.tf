resource "conductorone_integration_jira_cloud" "jira_cloud" {
  app_id = conductorone_app.jira_cloud.id
  user_ids = [
    conductorone_user.admin.id
  ]
  jiracloud_domain                    = "..."
  jiracloud_username                  = "..."
  jiracloud_apikey                    = "..."
  enable_external_ticket_provisioning = false
}
