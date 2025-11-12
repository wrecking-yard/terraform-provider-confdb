terraform {
  required_providers {
    confdb = {
      source = "localhost/dev/confdb"
    }
  }
}

provider "confdb" {
  region       = "northeurope"
  subscription = "sub1"
  environment  = "lab"
}

data "confdb_default_vnet" "default_vnet" {}

output "default_vnet" {
  value = data.confdb_default_vnet.default_vnet
}
