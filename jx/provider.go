package jx

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	p := createProvider()
	p.ConfigureFunc = providerConfigure(p)
	return p
}

func createProvider() *schema.Provider {
	p := &schema.Provider{
		Schema: map[string]*schema.Schema{},
	}
	return p
}

var descriptions map[string]string

func init() {
	descriptions = map[string]string{
		/*		"token": "The OAuth token used to connect to GitHub.",

				"organization": "The GitHub organization name to manage.",

				"base_url": "The GitHub Base API URL",
		*/}
}

func providerConfigure(p *schema.Provider) schema.ConfigureFunc {
	return func(d *schema.ResourceData) (interface{}, error) {
		options := &TerraformOptions{}
		return options, nil
	}
}
