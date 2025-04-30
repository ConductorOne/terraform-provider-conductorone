resource "conductorone_integration_jira_cloud_v2" "jira_cloud_v2" {
  app_id = conductorone_app.jira_cloud_v2.id
  user_ids = [
    conductorone_user.admin.id
  ]
  jiracloud_domain                    = "..."
  jiracloud_username                  = "..."
  jiracloud_apikey                    = "..."
  enable_external_ticket_provisioning = false
  jiracloud_project_keys              = ["..."]
  jiracloud_skip_project_participants = false
}
