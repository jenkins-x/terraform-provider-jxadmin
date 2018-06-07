package jx

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceEnvironment() *schema.Resource {
	return &schema.Resource{
		Create: resourceEnvironmentCreate,
		Read:   resourceEnvironmentRead,
		Update: resourceEnvironmentUpdate,
		Delete: resourceEnvironmentDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"promotion_strategy": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"order": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"namespace": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceEnvironmentCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceEnvironmentRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceEnvironmentUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceEnvironmentDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
