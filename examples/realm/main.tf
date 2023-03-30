terraform {
  required_providers {
    accessmanager = {
      source = "github.com/jralmaraz/forgerock-am"
    }
  }
}

variable "realm_name" {
  type    = string
  default = "/"
}

data "accessmanager_realms" "all" {}

# Returns all realms
output "all_realms" {
  value = data.accessmanager_realms.all.realms
}

output "realm" {
  value = {
    for realm in data.accessmanager_realms.all.realms :
    realm._id => realm
    if realm.name == var.realm_name
  }
}
