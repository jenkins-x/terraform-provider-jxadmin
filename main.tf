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
  config_path = "/tmp/bf8e5b76ce425c55d740fef2140d225a/config"
//  name = "my-jx-cluster"
//  endpoint = "https://127.0.0.1"
//  certificate = "some generated cert..."
}

resource "jxadmin_operator" "foo" {
  bot_user = "bot1"
  bot_token = "abc123"
}