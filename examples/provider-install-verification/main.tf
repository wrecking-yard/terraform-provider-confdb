terraform {
  required_providers {
    confdb = {
      source = "localhost/dev/confdb"
    }
  }
}

provider "confdb" {
  region = "westeurope"
  environment = "dev"
  subscription = "f9ijewc9ijiwemkcd"
}

data "confdb_vnet" "some_vnet" {}
data "confdb_default_vnet" "default_vnet" {}

output "x" {
  value = data.confdb_vnet.some_vnet
}

output "y" {
  value = data.confdb_default_vnet.default_vnet
}
