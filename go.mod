module github.com/jenkins-x/terraform-provider-jxadmin

go 1.12

require (
	github.com/gregjones/httpcache v0.0.0-20190611155906-901d90724c79 // indirect
	github.com/hashicorp/terraform v0.13.2
	github.com/jenkins-x/jx-api v0.0.17 // indirect
	github.com/mitchellh/go-homedir v1.1.0
	k8s.io/apimachinery v0.16.12
	k8s.io/client-go v10.0.0+incompatible
	k8s.io/kube-aggregator v0.0.0-20191025230902-aa872b06629d
)

// Override invalid go-autorest pseudo-version. This can be removed once
// all transitive dependencies on go-autorest use correct pseudo-versions.
// See https://tip.golang.org/doc/go1.13#version-validation
// and https://github.com/Azure/go-autorest/issues/481
replace (
	github.com/Azure/go-autorest v11.1.2+incompatible => github.com/Azure/go-autorest v12.1.0+incompatible
	k8s.io/client-go => k8s.io/client-go v0.0.0-20190918160344-1fbdaa4c8d90
)
