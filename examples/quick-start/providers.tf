# include the required provider
terraform {
  required_providers {
    definednet = {
      source = "registry.terraform.io/rfletcher/definednet"
    }
  }
}

# configure the provider
provider "definednet" {
  api_key = var.definednet_api_key
}
