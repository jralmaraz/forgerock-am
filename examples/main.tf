terraform {
  required_providers {
    accessmanager = {
      version = "0.1"
      source = "github.com/jralmaraz/forgerock-am"
    }
  }
}

provider "accessmanager" {
  username =  "amadmin"
  password = "TJUQzLoKOWHS937lEp3rWXJs"
}

module "realm" {
  source = "./realm"

  realm_name = "/"
}

output "realm" {
  value = module.realm.realm
}
