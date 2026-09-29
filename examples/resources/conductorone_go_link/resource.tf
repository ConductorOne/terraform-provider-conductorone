resource "conductorone_go_link" "my_golink" {
  description  = "...my_description..."
  display_name = "...my_display_name..."
  prerequisite = {
    app_entitlement_id = "...my_app_entitlement_id..."
    app_id             = "...my_app_id..."
  }
  routes = [
    {
      is_canonical   = true
      route_template = "...my_route_template..."
    }
  ]
  target = {
    action = {
      action_id = "...my_action_id..."
    }
    chooser = "{ \"see\": \"documentation\" }"
    redirect = {
      url_template = "...my_url_template..."
    }
  }
}