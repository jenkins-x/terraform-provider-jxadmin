package jx

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("JX_CLUSTER_NAME", nil),
			},
			"endpoint": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("JX_CLUSTER_ENDPOINT", nil),
			},
			"certificate": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("JX_CLUSTER_CERT", nil),
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
