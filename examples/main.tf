terraform {
  required_providers {
    accessmanager = {
      versions = ["0.1"]
      source = "github.com/jralmaraz/accessmanager"
    }
  }
}

provider "accessmanager" {
  xopenamusername = "amadmin"
  xopenampassword = "knzn2avs0peeyuc8hi7zwq4vd7ym0sec"
}

module "realm" {
  source = "./realm"

  realm_name = "/"
}

output "realm" {
  value = module.realm.realm
}
