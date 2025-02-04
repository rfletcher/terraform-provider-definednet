// include the provider
terraform {
  required_providers {
    definednet = {
      source = "registry.terraform.io/rfletcher/definednet"
    }
  }
}

provider "definednet" {
  // provide your Defined `api_key` here, or set
  // TF_DN_API_KEY in your environment
  api_key = null
}
