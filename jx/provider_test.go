package jx

import (
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/jenkins-x/jx/pkg/jx/cmd"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	// use test provider to use a fake cluster
	// 	testAccProvider = Provider().(*schema.Provider)
	testAccProvider = NewTestProvider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"jx": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	// use test provider to use a fake cluster
	// if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
	if err := NewTestProvider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	NewTestProvider()
}

func testAccPreCheck(t *testing.T) {
}

func NewTestProvider() terraform.ResourceProvider {
	p := createProvider()
	p.ConfigureFunc = testProviderConfigure(p)
	return p
}

func testProviderConfigure(p *schema.Provider) schema.ConfigureFunc {
	return func(d *schema.ResourceData) (interface{}, error) {
		options := &TerraformOptions{
			CommonOptions: newCommonOptions(),
			TestMode:      true,
		}
		commonOptions := &options.CommonOptions
		cmd.ConfigureTestOptionsWithResources(commonOptions,
			[]runtime.Object{
				&rbacv1.ClusterRole{
					ObjectMeta: metav1.ObjectMeta{
						Name: "cluster-admin",
					},
					Rules: []rbacv1.PolicyRule{
						{
							APIGroups: []string{"*"},
							Resources: []string{"*"},
							Verbs:     []string{"*"},
						},
						{
							NonResourceURLs: []string{"*"},
							Verbs:           []string{"*"},
						},
					},
				},
			},
			[]runtime.Object{})

		return options, nil
	}
}
