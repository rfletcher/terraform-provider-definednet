# include the required providers
terraform {
  required_providers {
    cloudflare = {
      source  = "cloudflare/cloudflare"
      version = "4.52.0"
    }

    definednet = {
      source = "registry.terraform.io/rfletcher/definednet"
    }
  }
}

# configure providers
provider "cloudflare" {
  api_token = var.cloudflare_api_token
}
provider "definednet" {
  api_key = var.definednet_api_key
}
