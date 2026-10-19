resource "conductorone_mcp_server" "my_mcp_server" {
  access_profile_ids = [
    "..."
  ]
  acknowledged_finding_ids = [
    "..."
  ]
  app_id = "...my_app_id..."
  app_managed_state_binding_ref = {
    app_id           = "...my_app_id..."
    resource_id      = "...my_resource_id..."
    resource_type_id = "...my_resource_type_id..."
  }
  data_sensitivity = "MCP_SERVER_DATA_SENSITIVITY_RESTRICTED"
  description      = "...my_description..."
  display_name     = "...my_display_name..."
  external_config = {
    basic_auth = {
      password = "...my_password..."
      username = "...my_username..."
    }
    bearer_token = {
      token = "...my_token..."
    }
    custom_header = {
      header_name  = "...my_header_name..."
      header_value = "...my_header_value..."
    }
    none = {
      # ...
    }
    oauth2 = {
      authorize_url  = "...my_authorize_url..."
      client_id      = "...my_client_id..."
      client_id_mode = "MCP_SERVER_AUTH_OAUTH2_CLIENT_ID_MODE_DCR"
      client_secret  = "...my_client_secret..."
      code_challenge_methods_supported = [
        "..."
      ]
      extra_authorize_params = {
        key = "value"
      }
      extra_token_params = {
        key = "value"
      }
      issuer_url      = "...my_issuer_url..."
      jwt_audience    = "...my_jwt_audience..."
      jwt_issuer      = "...my_jwt_issuer..."
      jwt_private_key = "...my_jwt_private_key..."
      jwt_subject     = "...my_jwt_subject..."
      mode            = "MCP_SERVER_AUTH_OAUTH2_MODE_SERVICE"
      pkce            = "...my_pkce..."
      scopes = [
        "..."
      ]
      scopes_supported = [
        "..."
      ]
      token_url = "...my_token_url..."
    }
    require_tool_approval = "OPTIONAL_BOOL_TRUE"
    token_sharing         = "MCP_SERVER_TOKEN_SHARING_PER_USER"
    transport_type        = "MCP_SERVER_TRANSPORT_TYPE_UNSPECIFIED"
    url                   = "...my_url..."
  }
  hosted_config = {
    aws_sigv4 = {
      access_key_id     = "...my_access_key_id..."
      secret_access_key = "...my_secret_access_key..."
      session_token     = "...my_session_token..."
    }
    basic_auth = {
      password = "...my_password..."
      username = "...my_username..."
    }
    bearer_token = {
      token = "...my_token..."
    }
    config_fields = {
      key = "value"
    }
    custom_header = {
      header_name  = "...my_header_name..."
      header_value = "...my_header_value..."
    }
    google_service_account = {
      credentials_json = "...my_credentials_json..."
      scopes = [
        "..."
      ]
    }
    mcp_server_catalog_id = "...my_mcp_server_catalog_id..."
    none = {
      # ...
    }
    oauth2 = {
      authorize_url  = "...my_authorize_url..."
      client_id      = "...my_client_id..."
      client_id_mode = "MCP_SERVER_AUTH_OAUTH2_CLIENT_ID_MODE_UNSPECIFIED"
      client_secret  = "...my_client_secret..."
      code_challenge_methods_supported = [
        "..."
      ]
      extra_authorize_params = {
        key = "value"
      }
      extra_token_params = {
        key = "value"
      }
      issuer_url      = "...my_issuer_url..."
      jwt_audience    = "...my_jwt_audience..."
      jwt_issuer      = "...my_jwt_issuer..."
      jwt_private_key = "...my_jwt_private_key..."
      jwt_subject     = "...my_jwt_subject..."
      mode            = "MCP_SERVER_AUTH_OAUTH2_MODE_SERVICE"
      pkce            = "...my_pkce..."
      scopes = [
        "..."
      ]
      scopes_supported = [
        "..."
      ]
      token_url = "...my_token_url..."
    }
    require_tool_approval = "OPTIONAL_BOOL_UNSPECIFIED"
    source_app_id         = "...my_source_app_id..."
    token_sharing         = "MCP_SERVER_TOKEN_SHARING_UNSPECIFIED"
  }
  mcp_server_service_delete_request = {
    # ...
  }
  oauth_diagnostic_id = "...my_oauth_diagnostic_id..."
  server_type         = "MCP_SERVER_TYPE_EXTERNAL"
  tool_prefix         = "...my_tool_prefix..."
  tunnel_appliance_id = "...my_tunnel_appliance_id..."
  tunnel_path         = "...my_tunnel_path..."
  tunnel_service_name = "...my_tunnel_service_name..."
  tunneled            = false
}