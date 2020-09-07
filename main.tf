terraform {
  required_providers {
    kubernetes = {
      source = "kubernetes"
    }

    // used for developing locally, need to add terraform CLI config for provider_installation to change location of plugin to local
    jxadmin = {
      source = "github.com/jenkins-x/jxadmin"
      version = "9.9.9"
    }
  }
  required_version = ">= 0.13"
}

provider "jxadmin" {
  name = "my-jx-cluster"
  endpoint = "https://127.0.0.1"
  certificate = "some generated cert..."
}

resource "jxadmin_operator" "foo" {
  pipelineUser = "bot1"
}