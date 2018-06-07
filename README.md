A [Jenkins X](https://jenkins-x.io/) provider for [Terraform](https://www.terraform.io/)

<a href="http://jenkins-x.io/">
  <img src="http://jenkins-x.io/img/profile.png" alt="Jenkins X icon" width="100" height="123"/>
</a>

This is a work in progress and should definitely not be considered stable!

# To build
```
mkdir -p $GOPATH/src/github.com/jenkins-x
cd $GOPATH/src/github.com/jenkins-x
git clone https://github.com/jenkins-x/terraform-provider-jx
cd terraform-provider-jx
make build
mkdir -p ~/.terraform.d/plugins/darwin_amd64
ln -s $GOPATH/bin/terraform-provider-jx ~/.terraform.d/plugins/darwin_amd64/terraform-provider-jx
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
Refreshing Terraform state in-memory prior to plan...
The refreshed state will be used to calculate this plan, but will not be
persisted to local or remote state storage.


------------------------------------------------------------------------

An execution plan has been generated and is shown below.
Resource actions are indicated with the following symbols:
  + create

Terraform will perform the following actions:

  + jx_environment.production
      id:                 <computed>
      name:               "production"
      namespace:          "jx-production"
      order:              "200"
      promotion_strategy: "manual"

  + jx_environment.staging
      id:                 <computed>
      name:               "staging"
      namespace:          "jx-staging"
      order:              "100"
      promotion_strategy: "auto"

  + jx_team.team1
      id:                 <computed>
      name:               "team1"

  + jx_team.team2
      id:                 <computed>
      name:               "team2"


Plan: 4 to add, 0 to change, 0 to destroy.

------------------------------------------------------------------------

Note: You didn't specify an "-out" parameter to save this plan, so Terraform
can't guarantee that exactly these actions will be performed if
"terraform apply" is subsequently run.
```