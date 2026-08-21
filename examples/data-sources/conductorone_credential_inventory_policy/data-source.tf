data "conductorone_credential_inventory_policy" "my_credentialinventorypolicy" {
  page_size  = 5
  page_token = "...my_page_token..."
  query      = "...my_query..."
  refs = [
    {
      id = "...my_id..."
    }
  ]
}