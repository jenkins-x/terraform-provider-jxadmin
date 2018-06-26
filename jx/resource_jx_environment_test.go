package jx

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/jenkins-x/jx/pkg/apis/jenkins.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestAccJXEnvironment_basic(t *testing.T) {
	var env v1.Environment
	randString := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	firstNamespace := fmt.Sprintf("tf-acc-namespace-%s", randString)
	updatedNamespace := fmt.Sprintf("tf-acc-namespace-updated-%s", randString)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckJXEnvironmentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccJXEnvironmentConfig(randString),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckJXEnvironmentExists("jx_environment.foo", &env),
					testAccCheckJXEnvironmentAttributes(&env, "foo", firstNamespace),
				),
			},
			{
				Config: testAccJXEnvironmentUpdateConfig(randString),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckJXEnvironmentExists("jx_environment.foo", &env),
					testAccCheckJXEnvironmentAttributes(&env, "foo", updatedNamespace),
				),
			},
		},
	})
}

func TestAccJXEnvironment_importBasic(t *testing.T) {
	randString := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckJXEnvironmentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccJXEnvironmentConfig(randString),
			},
			{
				ResourceName:      "jx_environment.foo",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckJXEnvironmentExists(n string, env *v1.Environment) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not Found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Environment ID is set")
		}
		options := testAccProvider.Meta().(*TerraformOptions)
		client, ns, err := options.JXClient()
		if err != nil {
			return fmt.Errorf("Could not create Jenkins X client: %s", err)
		}

		namespace, name := toNamespaceAndName(rs.Primary.ID, ns)
		loadedEnv, err := client.JenkinsV1().Environments(namespace).Get(name, metav1.GetOptions{})
		if err != nil {
			return err
		}
		*env = *loadedEnv
		return nil
	}
}

func testAccCheckJXEnvironmentAttributes(env *v1.Environment, name string, namespace string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if env.Name != name {
			return fmt.Errorf("Environment name does not match: %s, %s", env.Name, name)
		}
		specNs := env.Spec.Namespace
		if specNs != namespace {
			return fmt.Errorf("Environment namespace does not match: %s, %s", specNs, namespace)
		}
		return nil
	}
}

func testAccCheckJXEnvironmentDestroy(s *terraform.State) error {
	options := testAccProvider.Meta().(*TerraformOptions)
	client, ns, err := options.JXClient()
	if err != nil {
		return fmt.Errorf("Could not create Jenkins X client: %s", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "jx_environment" {
			continue
		}

		env, err := client.JenkinsV1().Environments(ns).Get(toResourceName(ns, rs.Primary.ID), metav1.GetOptions{})
		if err == nil {
			if env != nil && env.Name == rs.Primary.ID {
				return fmt.Errorf("Environment still exists")
			}
		}
		// TODO filter out does not exist errors
		/*		if resp.StatusCode != 404 {
					return err
				}
		*/
		return nil
	}
	return nil
}

func testAccJXEnvironmentConfig(randString string) string {
	return fmt.Sprintf(`
resource "jx_environment" "foo" {
	name = "foo"
	namespace = "tf-acc-namespace-%s"
	kind = "Permanent"
	order = 123
}
`, randString)
}

func testAccJXEnvironmentUpdateConfig(randString string) string {
	return fmt.Sprintf(`
resource "jx_environment" "foo" {
	name = "foo"
	namespace = "tf-acc-namespace-updated-%s"
	kind = "Permanent"
	order = 456
}
`, randString)
}
