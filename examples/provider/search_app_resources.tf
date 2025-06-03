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
      app_id = "2xVg4U1nyh7SRqjC5FT10fqT5Ju"
      id     = "287oY0rG4UirjDNFEYguMBvxyim"
    },
    {
      app_id = "2xVg4U1nyh7SRqjC5FT10fqT5Ju"
      id     = "2xpUq7cnIRgFyoV752cxj5jBwId"
    },
    {
      app_id = "2xVg4U1nyh7SRqjC5FT10fqT5Ju"
      id     = "2xpUqDccUIl7hoci6IEeCT2Aw8O"
    },
  ]
}