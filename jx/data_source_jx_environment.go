package jx

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jenkins-x/jx/pkg/apis/jenkins.io/v1"
	"github.com/jenkins-x/jx/pkg/client/clientset/versioned"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	environmentSchema = map[string]*schema.Schema{
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"kind": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "Permanent",
		},
		"label": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"namespace": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"order": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		"cluster": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"promotion_strategy": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "Auto",
		},
		"source_kind": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"git_url": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"git_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
	}
)

func dataSourceJXEnvironment() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceJXEnvironmentRead,

		Schema: environmentSchema,
	}
}

func dataSourceJXEnvironmentRead(d *schema.ResourceData, meta interface{}) error {
	name := d.Get("name").(string)
	log.Printf("[INFO] Refreshing JX Environment: %s", name)

	options := meta.(*TerraformOptions)
	client, ns, err := options.JXClient()
	if err != nil {
		return fmt.Errorf("Could not create Jenkins X client: %s", err)
	}

	env, err := getJXEnvironmentByName(client, ns, name)
	if err != nil {
		return err
	}

	loadEnvironmentResource(d, ns, env)
	return nil
}

func loadEnvironmentResource(d *schema.ResourceData, ns string, env *v1.Environment) {
	spec := &env.Spec
	d.SetId(ns + "/" + env.Name)
	d.Set("name", env.Name)
	d.Set("kind", string(spec.Kind))
	d.Set("label", spec.Label)
	d.Set("namespace", string(spec.Namespace))
	d.Set("order", int(spec.Order))
	d.Set("cluster", spec.Cluster)
	d.Set("promotion_strategy", string(spec.PromotionStrategy))
	d.Set("source_kind", string(spec.Source.Kind))
	d.Set("git_url", spec.Source.URL)
	d.Set("git_ref", spec.Source.Ref)
}

func getJXEnvironmentByName(jxClient *versioned.Clientset, ns string, name string) (*v1.Environment, error) {
	env, err := jxClient.JenkinsV1().Environments(ns).Get(name, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("Could not find Environment in namespace %s with name: %s: %s", ns, name, err)
	}
	return env, nil

}
