A [Jenkins X](https://jenkins-x.io/) provider for [Terraform](https://www.terraform.io/)

This is a work in progress and should definitely not be considered stable!

# To build
```
mkdir -p $GOPATH/src/github.com/jenkins-x
cd $GOPATH/src/github.com/jenkins-x
git clone https://github.com/jenkins-x/terraform-provider-jx-admin
cd terraform-provider-jx-admin
make build
mkdir -p ~/.terraform.d/plugins/darwin_amd64
ln -s $GOPATH/bin/terraform-provider-jx-admin ~/.terraform.d/plugins/darwin_amd64/terraform-provider-jx-admin
``` 

To check that the provider is installed correctly

```
terraform init
```
 
You should see an output like...

```
$ terraform init

Initializing provider plugins...

Terraform has been successfully initialized!

You may now begin working with Terraform. Try running "terraform plan" to see
any changes that are required for your infrastructure. All Terraform commands
should now work.

If you ever set or change modules or backend configuration for Terraform,
rerun this command to reinitialize your working directory. If you forget, other
commands will detect it and remind you to do so if necessary.
```

Check that the plugin works by running a terraform plan

```
$ terraform plan

```