package jxadmin

import (
	"errors"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func resourceJxAdminOperator() *schema.Resource {
	return &schema.Resource{
		Create: resourceJxAdminOperatorCreate,
		Read:   resourceJxAdminOperatorRead,
		Exists: resourceJxAdminOperatorExists,
		Update: resourceJxAdminOperatorUpdate,
		Delete: resourceJxAdminOperatorDelete,
		//Importer: &schema.ResourceImporter{
		//	State: schema.ImportStatePassthrough,
		//},

		Schema: map[string]*schema.Schema{
			"bot_user": {
				Type:        schema.TypeString,
				Description: "Bot username used to authenticate with Git provider",
				Required:    true,
			},
			"bot_token": {
				Type:        schema.TypeString,
				Description: "Bot token used to authenticate with Git provider",
				Required:    true,
			},
		},
	}
}

func resourceJxAdminOperatorCreate(d *schema.ResourceData, meta interface{}) error {
	conn, err := meta.(KubeClientsets).MainClientset()
	if err != nil {
		return err
	}
	namespaces, err := conn.CoreV1().Namespaces().List(v1.ListOptions{})
	if err != nil {
		return err
	}
	log.Printf("1")
	if namespaces == nil {
		return errors.New("no namespaces found")
	}
	log.Printf("2")
	for _, ns := range namespaces.Items {
		log.Printf("%s", ns.Name)
	}
	log.Printf("About to check JX3 resource exists")
	d.SetId("123abc")
	return resourceJxAdminOperatorRead(d, meta)
}

func resourceJxAdminOperatorExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	_, err := meta.(KubeClientsets).MainClientset()
	if err != nil {
		return false, err
	}

	return true, nil
}
func resourceJxAdminOperatorUpdate(d *schema.ResourceData, meta interface{}) error {
	_, err := meta.(KubeClientsets).MainClientset()
	if err != nil {
		return err
	}
	log.Printf("About to update JX3 resource")
	return nil
}

func resourceJxAdminOperatorRead(d *schema.ResourceData, meta interface{}) error {
	_, err := meta.(KubeClientsets).MainClientset()
	if err != nil {
		return err
	}
	log.Printf("About to read JX3 resource")
	return nil
}

func resourceJxAdminOperatorDelete(d *schema.ResourceData, meta interface{}) error {
	_, err := meta.(KubeClientsets).MainClientset()
	if err != nil {
		return err
	}
	log.Printf("About to delete JX3 resource")
	return nil
}
