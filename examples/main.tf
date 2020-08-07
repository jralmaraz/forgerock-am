terraform {
  required_providers {
    accessmanager = {
      versions = ["0.1"]
      source = "github.com/jralmaraz/accessmanager"
    }
  }
}

provider "accessmanager" {
  xopenamusername = ""
  xopenampassword = ""
}

module "realm" {
  source = "./realm"

  realm_name = "/"
}

output "realm" {
  value = module.realm.realm
}
