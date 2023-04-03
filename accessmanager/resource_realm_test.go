package main

import (
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccResourceRealmCreate(t *testing.T) {
	realmName := "test"
	realmAliases := []string{"dev.example.com", "am-config", "am"}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckRealmDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
                    resource "realm" "test" {
                        name = "%s"
                        aliases = %v
                    }
                `, realmName, realmAliases),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRealmExists("realm.test"),
					resource.TestCheckResourceAttr("realm.test", "name", realmName),
					resource.TestCheckResourceAttr("realm.test", "aliases.#", fmt.Sprintf("%d", len(realmAliases))),
					resource.TestCheckResourceAttr("realm.test", "aliases.0", realmAliases[0]),
					resource.TestCheckResourceAttr("realm.test", "aliases.1", realmAliases[1]),
					resource.TestCheckResourceAttr("realm.test", "aliases.2", realmAliases[2]),
				),
			},
		},
	})
}

func testAccCheckRealmExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("resource not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no ID is set for resource: %s", n)
		}

		return nil
	}
}

func testAccCheckRealmDestroy(s *terraform.State) error {
	// Here, you can implement a function that checks if the resource was actually deleted
	// In this case, you can check if the realm was deleted by querying the ForgeRock Access Manager API
	return nil
}

func testAccPreCheck(t *testing.T) {
	// Here, you can implement a function that checks if the environment variables required for testing are set
	// In this case, you can check if the ForgeRock Access Manager API endpoint, username and password are set
}

var testAccProviders = map[string]terraform.ResourceProvider{
	"forgerock": Provider(),
}
