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
  environment  = "dev"
}

data "confdb_default_vnet" "default_vnet" {}
data "confdb_default_subnet" "default_subnet" {
  vnet_name = data.confdb_default_vnet.default_vnet.name
}

output "default_vnet" {
  value = data.confdb_default_vnet.default_vnet
}

output "default_subnet" {
  value = data.confdb_default_subnet.default_subnet
}
