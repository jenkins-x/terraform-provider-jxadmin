package jx

import (
	"os"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/jenkins-x/jx/pkg/jx/cmd"
	cmdutil "github.com/jenkins-x/jx/pkg/jx/cmd/util"
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
		/*
			// TODO

			"token": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("GITHUB_TOKEN", nil),
				Description: descriptions["token"],
			},
			"organization": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("GITHUB_ORGANIZATION", nil),
				Description: descriptions["organization"],
			},
			"base_url": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("GITHUB_BASE_URL", ""),
				Description: descriptions["base_url"],
			},
		*/

		ResourcesMap: map[string]*schema.Resource{
			"jx_environment": resourceEnvironment(),
		},

		DataSourcesMap: map[string]*schema.Resource{
			"jx_environment": dataSourceJXEnvironment(),
		},
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
		options := &TerraformOptions{
			CommonOptions: cmd.CommonOptions{
				Factory:   cmdutil.NewFactory(),
				Out:       os.Stdout,
				Err:       os.Stderr,
				BatchMode: true,
			},
		}
		return options, nil
	}
}
