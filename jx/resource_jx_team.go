package jx

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jenkins-x/jx/pkg/jx/cmd"
)

var (
	teamSchema = map[string]*schema.Schema{
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"version": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"kube_kind": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "kubernetes",
		},
		"domain": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"helm3": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"helm_tls": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
	}
)

func resourceTeam() *schema.Resource {
	return &schema.Resource{
		Create: resourceTeamCreate,
		Read:   resourceTeamRead,
		Update: resourceTeamUpdate,
		Delete: resourceTeamDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: teamSchema,
	}
}

func resourceTeamRead(d *schema.ResourceData, meta interface{}) error {
	// TODO
	return nil
}

func resourceTeamCreate(d *schema.ResourceData, meta interface{}) error {
	return resourceTeamUpsert(d, meta, true)
}

func resourceTeamUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceTeamUpsert(d, meta, false)
}

func resourceTeamUpsert(d *schema.ResourceData, meta interface{}, isCreate bool) error {
	options := meta.(*TerraformOptions)

	installOptions := &cmd.InstallOptions{}
	commonOptions := options.CommonOptions
	installOptions.CommonOptions = commonOptions
	installOptions.InitOptions.CommonOptions = commonOptions
	installOptions.CreateJenkinsUserOptions.CommonOptions = commonOptions

	installOptions.BatchMode = true
	installOptions.Headless = true

	flags := &installOptions.Flags
	flags.Namespace = asString(d.Get("name"))
	flags.Provider = asString(d.Get("kube_kind"))
	flags.Domain = asString(d.Get("domain"))
	flags.HelmTLS = asBool(d.Get("helm_tls"))

	// TODO
	//flags.Version = asBool(d.Get("version"))

	initOptions := &installOptions.InitOptions.Flags

	fmt.Printf("TestMode %v\n", options.TestMode)
	if options.TestMode {
		initOptions.HelmBin = "echo"
	} else {
		initOptions.Helm3 = asBool(d.Get("helm_tls"))
	}

	fmt.Printf("helm bin is %s\n", installOptions.InitOptions.HelmBinary())
	return installOptions.Run()
}

func resourceTeamDelete(d *schema.ResourceData, meta interface{}) error {
	// TODO invoke jx uninstall in the teams namespace
	return fmt.Errorf("TODO")
}
