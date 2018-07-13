package jx

import (
	"strings"

	"github.com/jenkins-x/jx/pkg/jx/cmd"
)

type TerraformOptions struct {
	cmd.CommonOptions

	TestMode bool
}

func toResourceName(ns string, name string) string {
	prefix := ns + "/"
	return strings.TrimPrefix(name, prefix)
}

// toNamespaceAndName returns the namespace and name for the given resource ID
func toNamespaceAndName(id string, defaultNamespace string) (string, string) {
	values := strings.Split(id, "/")
	if len(values) < 2 {
		return defaultNamespace, id
	}
	return values[0], values[1]
}
