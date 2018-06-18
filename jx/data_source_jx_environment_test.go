package jx

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccJXEnvironmentSource_noMatchReturnsError(t *testing.T) {
	slug := "non-existing"
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckJXEnvironmentDataSourceConfig(slug),
				ExpectError: regexp.MustCompile(`.*Could not find Environment.*`),
			},
		},
	})
}

func testAccCheckJXEnvironmentDataSourceConfig(name string) string {
	return fmt.Sprintf(`
data "jx_environment" "something" {
	name = "%s"
}
`, name)
}
