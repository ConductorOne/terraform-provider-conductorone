resource "conductorone_mcp_server" "hosted" {
  app_id       = "...my_app_id..."
  display_name = "Hosted MCP server"
  server_type  = "MCP_SERVER_TYPE_HOSTED"

  hosted_config = {
    mcp_server_catalog_id = "...my_mcp_server_catalog_id..."
    config_fields = {
      key = "value"
    }
    token_sharing         = "MCP_SERVER_TOKEN_SHARING_SHARED"
    require_tool_approval = "OPTIONAL_BOOL_TRUE"

    bearer_token = {
      token = "...my_token..."
    }
  }
}

resource "conductorone_mcp_server" "external" {
  app_id       = "...my_app_id..."
  display_name = "External MCP server"
  server_type  = "MCP_SERVER_TYPE_EXTERNAL"

  external_config = {
    url                   = "https://mcp.example.com/mcp"
    transport_type        = "MCP_SERVER_TRANSPORT_TYPE_STREAMABLE_HTTP"
    token_sharing         = "MCP_SERVER_TOKEN_SHARING_SHARED"
    require_tool_approval = "OPTIONAL_BOOL_TRUE"

    bearer_token = {
      token = "...my_token..."
    }
  }
}
