package jx

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type teamSettings struct {
	Name string
}

func TestAccJXTeam_basic(t *testing.T) {
	var env teamSettings
	randString := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	firstNamespace := fmt.Sprintf("tf-acc-namespace-%s", randString)
	updatedNamespace := fmt.Sprintf("tf-acc-namespace-updated-%s", randString)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckJXTeamDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccJXTeamConfig(randString),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckJXTeamExists("jx_team.foo", &env),
					testAccCheckJXTeamAttributes(&env, "foo", firstNamespace),
				),
			},
			{
				Config: testAccJXTeamUpdateConfig(randString),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckJXTeamExists("jx_team.foo", &env),
					testAccCheckJXTeamAttributes(&env, "foo", updatedNamespace),
				),
			},
		},
	})
}

func TestAccJXTeam_importBasic(t *testing.T) {
	randString := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckJXTeamDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccJXTeamConfig(randString),
			},
			{
				ResourceName:      "jx_team.foo",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckJXTeamExists(n string, env *teamSettings) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not Found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Team ID is set")
		}
		options := testAccProvider.Meta().(*TerraformOptions)
		client, _, err := options.KubeClient()
		if err != nil {
			return fmt.Errorf("Could not create Kube client: %s", err)
		}

		name := rs.Primary.ID
		_, err = client.CoreV1().Namespaces().Get(name, metav1.GetOptions{})
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckJXTeamAttributes(env *teamSettings, name string, namespace string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if env.Name != name {
			return fmt.Errorf("Team name does not match: %s, %s", env.Name, name)
		}
		/*
			specNs := env.Spec.Namespace
			if specNs != namespace {
				return fmt.Errorf("Team namespace does not match: %s, %s", specNs, namespace)
			}
		*/return nil
	}
}

func testAccCheckJXTeamDestroy(s *terraform.State) error {
	options := testAccProvider.Meta().(*TerraformOptions)
	client, _, err := options.KubeClient()
	if err != nil {
		return fmt.Errorf("Could not create Jenkins X client: %s", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "jx_team" {
			continue
		}

		name := rs.Primary.ID
		_, err = client.CoreV1().Namespaces().Get(name, metav1.GetOptions{})
		if err == nil {
			return fmt.Errorf("Team still exists: %s", name)
		}
	}
	// TODO filter out does not exist errors
	/*		if resp.StatusCode != 404 {
				return err
			}
	*/
	return nil
}

func testAccJXTeamConfig(randString string) string {
	return fmt.Sprintf(`
resource "jx_team" "foo" {
	name = "%s"
}
`, randString)
}

func testAccJXTeamUpdateConfig(randString string) string {
	return fmt.Sprintf(`
resource "jx_team" "foo" {
	name = "%s"
}
`, randString)
}
