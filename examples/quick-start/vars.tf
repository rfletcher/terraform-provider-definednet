variable "definednet_api_key" {
  type = string
  description = "Your Defined Networking API key. Requires `hosts:list` scope."
  sensitive = true
}
