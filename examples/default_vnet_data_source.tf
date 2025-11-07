terraform {
  required_providers {
    confdb = {
      source = "localhost/dev/confdb"
    }
  }
}

provider "confdb" {
  region       = "westeurope"
  subscription = "subscription1"
  environment  = "dev"
}

data "confdb_default_vnet" "default_vnet" {}

output "default_vnet" {
  value = data.confdb_default_vnet.default_vnet
}
