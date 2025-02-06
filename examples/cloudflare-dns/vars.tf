variable "cloudflare_api_token" {
  type = string
  description = "Your Cloudflare API key. Requires DNS edit permission for the relevant zone."
  sensitive = true
}

variable "cloudflare_zone_name" {
  type = string
  description = "The name of the zone in which we'll create Nebula DNS records."
}

variable "definednet_api_key" {
  type = string
  description = "Your Defined Networking API key. Requires `hosts:list` scope."
  sensitive = true
}
