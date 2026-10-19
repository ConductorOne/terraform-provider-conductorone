package provider

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

func TestAccAccessProfileTypes(t *testing.T) {
	requireAccessProfileTypesAcceptance(t)

	for _, test := range []struct {
		name              string
		profileType       string
		published         bool
		requestBundle     bool
		visibleToEveryone bool
	}{
		{name: "requestable", profileType: "REQUEST_CATALOG_TYPE_CATALOG", published: true, visibleToEveryone: true},
		{name: "birthright", profileType: "REQUEST_CATALOG_TYPE_BUNDLE", visibleToEveryone: true},
		{name: "blended", profileType: "REQUEST_CATALOG_TYPE_CATALOG_AND_BUNDLE", published: true, requestBundle: true, visibleToEveryone: true},
	} {
		t.Run(test.name, func(t *testing.T) {
			resourceName := "conductorone_access_profile.test"
			resource.Test(t, resource.TestCase{
				ProtoV6ProviderFactories: testAccProviderFactories,
				Steps: []resource.TestStep{
					{
						Config: accessProfileTypeConfig(test.name, test.profileType, test.published, test.requestBundle, test.visibleToEveryone),
						Check: resource.ComposeAggregateTestCheckFunc(
							resource.TestCheckResourceAttr(resourceName, "type", test.profileType),
							resource.TestCheckResourceAttr(resourceName, "request_bundle", fmt.Sprint(test.requestBundle)),
							resource.TestCheckResourceAttr(resourceName, "visible_to_everyone", fmt.Sprint(test.visibleToEveryone)),
						),
					},
					{
						ResourceName:      resourceName,
						ImportState:       true,
						ImportStateVerify: true,
					},
				},
			})
		})
	}
}

func TestAccAccessProfileTypeRequiresReplacement(t *testing.T) {
	requireAccessProfileTypesAcceptance(t)

	const resourceName = "conductorone_access_profile.test"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: accessProfileTypeConfig("replace-type", "REQUEST_CATALOG_TYPE_CATALOG", true, false, false),
			},
			{
				Config: accessProfileTypeConfig("replace-type", "REQUEST_CATALOG_TYPE_BUNDLE", false, false, false),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionReplace),
					},
				},
				Check: resource.TestCheckResourceAttr(resourceName, "type", "REQUEST_CATALOG_TYPE_BUNDLE"),
			},
		},
	})
}

func TestAccBirthrightRejectsRequestableSettings(t *testing.T) {
	requireAccessProfileTypesAcceptance(t)

	for _, test := range []struct {
		name          string
		published     bool
		requestBundle bool
		errorPattern  string
	}{
		{
			name:         "published",
			published:    true,
			errorPattern: `access profile type Birthright does not support\s+entitlement request settings`,
		},
		{
			name:          "request bundle",
			requestBundle: true,
			errorPattern:  `access profile type Birthright does not support profile\s+membership requests`,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				ProtoV6ProviderFactories: testAccProviderFactories,
				Steps: []resource.TestStep{
					{
						Config:      accessProfileTypeConfig("invalid-birthright", "REQUEST_CATALOG_TYPE_BUNDLE", test.published, test.requestBundle, false),
						ExpectError: regexp.MustCompile(test.errorPattern),
					},
				},
			})
		})
	}
}

func TestAccAccessProfileVisibilityBindingsTypeCapabilities(t *testing.T) {
	requireAccessProfileTypesAcceptance(t)

	visibilityBinding := `
resource "conductorone_access_profile_visibility_bindings" "test" {
  catalog_id = conductorone_access_profile.test.id
  access_entitlements = [{
    app_id = conductorone_app.test.id
    id     = conductorone_custom_app_entitlement.test.id
  }]
}
`

	t.Run("requestable accepts visibility bindings", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: testAccProviderFactories,
			Steps: []resource.TestStep{
				{
					Config: accessProfileEntitlementConfig("visibility-requestable", "REQUEST_CATALOG_TYPE_CATALOG", visibilityBinding),
					Check: resource.TestCheckResourceAttrPair(
						"conductorone_access_profile_visibility_bindings.test", "catalog_id",
						"conductorone_access_profile.test", "id",
					),
				},
			},
		})
	})

	t.Run("birthright rejects visibility bindings", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: testAccProviderFactories,
			Steps: []resource.TestStep{
				{
					Config:      accessProfileEntitlementConfig("visibility-birthright", "REQUEST_CATALOG_TYPE_BUNDLE", visibilityBinding),
					ExpectError: regexp.MustCompile(`access profile type Birthright does not support\s+entitlement request settings`),
				},
			},
		})
	})
}

func TestAccBundleAutomationTypeCapabilities(t *testing.T) {
	requireAccessProfileTypesAcceptance(t)

	const resourceName = "conductorone_bundle_automation.test"

	t.Run("birthright accepts automation and corrective disable", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: testAccProviderFactories,
			Steps: []resource.TestStep{
				{
					Config: accessProfileEntitlementConfig("automation-birthright", "REQUEST_CATALOG_TYPE_BUNDLE", bundleAutomationConfig(true)),
					Check:  resource.TestCheckResourceAttr(resourceName, "enabled", "true"),
				},
				{
					Config: accessProfileEntitlementConfig("automation-birthright", "REQUEST_CATALOG_TYPE_BUNDLE", bundleAutomationConfig(false)),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
						},
					},
					Check: resource.TestCheckResourceAttr(resourceName, "enabled", "false"),
				},
			},
		})
	})

	t.Run("requestable rejects automation", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			ProtoV6ProviderFactories: testAccProviderFactories,
			Steps: []resource.TestStep{
				{
					Config:      accessProfileEntitlementConfig("automation-requestable", "REQUEST_CATALOG_TYPE_CATALOG", bundleAutomationConfig(true)),
					ExpectError: regexp.MustCompile(`access profile type Requestable does not support\s+membership automation`),
				},
			},
		})
	})
}

func requireAccessProfileTypesAcceptance(t *testing.T) {
	t.Helper()

	if os.Getenv("CONDUCTORONE_ACCESS_PROFILE_TYPES_ENABLED") != "1" {
		t.Skip("set CONDUCTORONE_ACCESS_PROFILE_TYPES_ENABLED=1 for a tenant with the access profile types feature")
	}
}

func accessProfileTypeConfig(name, profileType string, published, requestBundle, visibleToEveryone bool) string {
	return providerConfig + fmt.Sprintf(`
resource "conductorone_access_profile" "test" {
  display_name        = %q
  published           = %t
  request_bundle      = %t
  type                = %q
  visible_to_everyone = %t
}
`, "terraform-access-profile-type-"+name, published, requestBundle, profileType, visibleToEveryone)
}

func accessProfileEntitlementConfig(name, profileType, childResource string) string {
	published := profileType != "REQUEST_CATALOG_TYPE_BUNDLE"

	return providerConfig + fmt.Sprintf(`
resource "conductorone_app" "test" {
  display_name      = %q
  description       = "Access profile type acceptance fixture"
  identity_matching = "APP_USER_IDENTITY_MATCHING_STRICT"
}

resource "conductorone_app_resource_type" "test" {
  app_id        = conductorone_app.test.id
  display_name  = "GROUP"
  resource_type = "GROUP"
}

resource "conductorone_app_resource" "test" {
  app_id               = conductorone_app.test.id
  app_resource_type_id = conductorone_app_resource_type.test.id
  display_name         = "Access profile type fixture"
}

resource "conductorone_custom_app_entitlement" "test" {
  alias                = "tf_access_profile_type_member"
  app_id               = conductorone_app.test.id
  app_resource_id      = conductorone_app_resource.test.id
  app_resource_type_id = conductorone_app_resource_type.test.id
  display_name         = "Member"
  duration_grant       = "3601s"
  purpose              = "APP_ENTITLEMENT_PURPOSE_VALUE_ASSIGNMENT"
  slug                 = "member"
}

resource "conductorone_access_profile" "test" {
  display_name        = %q
  published           = %t
  request_bundle      = false
  type                = %q
  visible_to_everyone = false
}

%s
`, "terraform-access-profile-type-fixture-"+name, "terraform-access-profile-type-"+name, published, profileType, childResource)
}

func bundleAutomationConfig(enabled bool) string {
	return fmt.Sprintf(`
resource "conductorone_bundle_automation" "test" {
  request_catalog_id = conductorone_access_profile.test.id
  create_tasks       = false
  enabled            = %t
  entitlements = {
    entitlement_refs = [{
      app_id = conductorone_app.test.id
      id     = conductorone_custom_app_entitlement.test.id
    }]
  }
}
`, enabled)
}
