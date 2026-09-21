resource "conductorone_mcp_source" "catalog" {
  app_id       = "...my_app_id..."
  catalog_id   = "...my_catalog_id..."
  display_name = "...my_catalog_source_name..."
  source_type  = "MCP_SERVER_TYPE_CATALOG"

  hosted_config = jsonencode({
    none = {}
  })
}

resource "conductorone_mcp_source" "external" {
  app_id       = "...my_app_id..."
  display_name = "...my_external_source_name..."
  external_url = "https://mcp.example.com/mcp"
  source_type  = "MCP_SERVER_TYPE_EXTERNAL"

  external_config = jsonencode({
    none = {}
  })
}
