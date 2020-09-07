package jxadmin

import (
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"jx": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
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
		options := &TerraformOptions{}
		return options, nil
	}
}
