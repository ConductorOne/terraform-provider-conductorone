resource "conductorone_access_profile" "my_access_profile" {
  description                       = "...my_description..."
  display_name                      = "...my_display_name..."
  enrollment_behavior               = "REQUEST_CATALOG_ENROLLMENT_BEHAVIOR_BYPASS_ENTITLEMENT_REQUEST_POLICY"
  published                         = false
  request_bundle                    = false
  unenrollment_behavior             = "REQUEST_CATALOG_UNENROLLMENT_BEHAVIOR_LEAVE_ACCESS_AS_IS"
  unenrollment_entitlement_behavior = "REQUEST_CATALOG_UNENROLLMENT_ENTITLEMENT_BEHAVIOR_ENFORCE"
  visible_to_everyone               = false
}