package jx

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("HCLOUD_TOKEN", nil),
				Description: "The API token to access the hetzner cloud.",
			},
			"endpoint": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("HCLOUD_ENDPOINT", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"jx_team":        resourceTeam(),
			"jx_environment": resourceEnvironment(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	//opts := []hcloud.ClientOption{
	//	hcloud.WithToken(d.Get("token").(string)),
	//}
	//if endpoint, ok := d.GetOk("endpoint"); ok {
	//	opts = append(opts, hcloud.WithEndpoint(endpoint.(string)))
	//}
	//return hcloud.NewClient(opts...), nil
	return nil, nil
}
