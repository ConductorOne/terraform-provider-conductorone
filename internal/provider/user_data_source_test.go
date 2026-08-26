package provider

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceUserProfile(t *testing.T) {
	email := os.Getenv("CONDUCTORONE_TEST_USER_EMAIL")
	if email == "" {
		t.Skip("CONDUCTORONE_TEST_USER_EMAIL not set; requires a dev-integration user with profile data")
	}
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + fmt.Sprintf(`
				data "conductorone_user" "profile_test" {
					email = %q
				}
				`, email),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.conductorone_user.profile_test", "id"),
					resource.TestCheckResourceAttrWith("data.conductorone_user.profile_test", "profile", func(value string) error {
						var profile map[string]interface{}
						if err := json.Unmarshal([]byte(value), &profile); err != nil {
							return fmt.Errorf("profile is not a JSON object: %w (value: %s)", err, value)
						}
						if len(profile) == 0 {
							return fmt.Errorf("profile is empty: %s", value)
						}
						return nil
					}),
				),
			},
		},
	})
}
