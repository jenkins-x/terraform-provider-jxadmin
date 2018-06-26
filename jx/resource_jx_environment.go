package jx

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jenkins-x/jx/pkg/apis/jenkins.io/v1"
	"github.com/jenkins-x/jx/pkg/client/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func resourceEnvironment() *schema.Resource {
	return &schema.Resource{
		Create: resourceEnvironmentCreate,
		Read:   resourceEnvironmentRead,
		Update: resourceEnvironmentUpdate,
		Delete: resourceEnvironmentDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: environmentSchema,
	}
}

func resourceEnvironmentCreate(d *schema.ResourceData, meta interface{}) error {
	options := meta.(*TerraformOptions)
	client, ns, err := options.JXClient()
	if err != nil {
		return fmt.Errorf("Could not create Jenkins X client: %s", err)
	}

	id := d.Id()
	ns2, n := toNamespaceAndName(id, ns)
	name := asString(d.Get("name"))
	if name == "" {
		name = n
	}
	namespace := asString(d.Get("namespace"))
	if namespace == "" {
		namespace = ns2
		if ns2 == "" {
			namespace = ns
		}
	}

	create := false
	env, err := client.JenkinsV1().Environments(ns).Get(name, metav1.GetOptions{})
	if err != nil {
		create = true
		env = &v1.Environment{
			ObjectMeta: metav1.ObjectMeta{
				Name: name,
			},
		}
	}
	updateEnvironmentFromResourceData(d, env)
	if create {
		env, err = client.JenkinsV1().Environments(ns).Create(env)
	} else {
		env, err = client.JenkinsV1().Environments(ns).Update(env)
	}
	if err != nil {
		return err
	}

	d.SetId(toResourceName(ns, env.Name))
	return resourceEnvironmentRead(d, meta)
}

func updateEnvironmentFromResourceData(d *schema.ResourceData, env *v1.Environment) error {
	if env == nil {
		return fmt.Errorf("No Envrionment specified!")
	}
	spec := &env.Spec
	spec.Order = asInt32(d.Get("order"))

	envNs := asString(d.Get("namespace"))
	spec.Namespace = envNs

	label := asString(d.Get("label"))
	spec.Label = label
	promotionStrategy := asString(d.Get("promotion_strategy"))
	spec.PromotionStrategy = v1.PromotionStrategyType(promotionStrategy)
	envKind := asString(d.Get("kind"))
	spec.Kind = v1.EnvironmentKindType(envKind)

	sourceKind := asString(d.Get("source_kind"))
	gitUrl := d.Get("git_url")
	gitRef := d.Get("git_ref")

	source := &spec.Source
	source.Kind = v1.EnvironmentRepositoryType(sourceKind)
	source.URL = asString(gitUrl)
	source.Ref = asString(gitRef)
	return nil
}

func asInt32(n interface{}) int32 {
	i, ok := n.(int32)
	if ok {
		return i
	}
	i2, ok := n.(int)
	if ok {
		return int32(i2)
	}
	return 0
}

func asString(n interface{}) string {
	s, ok := n.(string)
	if ok {
		return s
	}
	return ""
}

func resourceEnvironmentRead(d *schema.ResourceData, meta interface{}) error {
	options := meta.(*TerraformOptions)
	client, ns, err := options.JXClient()
	if err != nil {
		return fmt.Errorf("Could not create Jenkins X client: %s", err)
	}

	env, err := getEnvironment(d, client, ns)
	if err != nil {
		d.SetId("")
		return nil
	}
	loadEnvironmentResource(d, ns, env)
	return nil
}

func resourceEnvironmentUpdate(d *schema.ResourceData, meta interface{}) error {
	options := meta.(*TerraformOptions)
	client, ns, err := options.JXClient()
	if err != nil {
		return fmt.Errorf("Could not create Jenkins X client: %s", err)
	}
	env, err := getEnvironment(d, client, ns)
	if err != nil {
		d.SetId("")
		return nil
	}

	updateEnvironmentFromResourceData(d, env)

	env, err = client.JenkinsV1().Environments(ns).Update(env)
	if err != nil {
		return err
	}

	d.SetId(toResourceName(ns, env.Name))
	return resourceEnvironmentRead(d, meta)
}

func resourceEnvironmentDelete(d *schema.ResourceData, meta interface{}) error {
	options := meta.(*TerraformOptions)
	client, ns, err := options.JXClient()
	if err != nil {
		return fmt.Errorf("Could not create Jenkins X client: %s", err)
	}
	name := toResourceName(ns, d.Id())
	return client.JenkinsV1().Environments(ns).Delete(name, &metav1.DeleteOptions{})
}

func getEnvironment(d *schema.ResourceData, client versioned.Interface, ns string) (*v1.Environment, error) {
	namespace, name := toNamespaceAndName(d.Id(), ns)
	env, err := client.JenkinsV1().Environments(namespace).Get(name, metav1.GetOptions{})
	return env, err
}
