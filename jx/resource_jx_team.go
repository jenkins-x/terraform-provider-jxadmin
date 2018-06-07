package jx

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceTeam() *schema.Resource {
	return &schema.Resource{
		Create: resourceTeamCreate,
		Read:   resourceTeamRead,
		Update: resourceTeamUpdate,
		Delete: resourceTeamDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceTeamCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceTeamRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceTeamUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceTeamDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
