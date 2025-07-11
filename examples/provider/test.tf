resource "conductorone_access_profile" "test_catalog" {
  display_name        = "catalog"
  description         = "terraform test"
  visible_to_everyone = "true"
  published           = "true"
}

resource "conductorone_access_profile_requestable_entries" "test_entries" {
  catalog_id = conductorone_access_profile.test_catalog.id
  app_entitlements = [
    {
      app_id = "2yh8YGCnRUP92bTwHj3yfzmauYJ"
      id     = "2yh8yQD5qE29UVwqfAFLMQsvm9s"
    },
    {
      app_id = "2yT54TrMaVkrchHuEJ4gAOW5w6x"
      id     = "2yT6BGbj08iKn8hHaplpl2grQUQ"
    },
    {
      app_id = "2yh8YGCnRUP92bTwHj3yfzmauYJ"
      id     = "2yh8yOzZyDIj2Dwq3MuOwP18ibW"
    }
  ]
}