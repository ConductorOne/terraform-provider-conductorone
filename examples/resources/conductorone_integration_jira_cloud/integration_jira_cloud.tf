resource "conductorone_integration_jira_cloud" "jira_cloud" {
  app_id             = conductorone_app.jira_cloud.id
  jiracloud_domain   = "..."
  jiracloud_username = "..."
  jiracloud_apikey   = "..."
}
