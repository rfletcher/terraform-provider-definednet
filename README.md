# terraform-provider-definednet

An unofficial [Terraform](https://www.terraform.io) provider for the
[Defined Networking API](https://docs.defined.net/api/defined-networking-api/).

> [!WARNING]
> This is functional, but very incomplete!

# Quick Start

This provider is only available locally, through Terraform's
[development overrides](https://developer.hashicorp.com/terraform/cli/config/config-file#development-overrides-for-provider-developers):

1. Install [Terraform](https://developer.hashicorp.com/terraform/install) and
   [Go](https://go.dev/doc/install)
2. Clone this repository
3. Make the plugin available to Terraform by editing `$HOME/.terraformrc` to include
   this configuration (be sure to use your actual *absolute* `$GOBIN` path --
   often `$HOME/go/bin`):

        provider_installation {
          dev_overrides {
            "registry.terraform.io/rfletcher/definednet" = "/path/to/go/bin"
          }
          direct {}
        }

4. Run `make` to build and install
5. Use the provider!

# Example

A simple example, [listing your hosts](https://docs.defined.net/api/hosts-list/):

```terraform
// include and configure the provider
terraform {
  required_providers {
    definednet = {
      source = "registry.terraform.io/rfletcher/definednet"
    }
  }
}
provider "definednet" {
  // configure `api_key`, or set TF_DN_API_KEY in your environment
  api_key = "dnkey-abc123"
}


// list hosts from the Defined API
data "definednet_hosts" "all" {}


// expose the data you're interested in
output "first_host" {
  value = data.definednet_hosts.all.hosts[0]
}
output "lighthouse_ips" {
  value = [for host in data.definednet_hosts.all.hosts : host.ip_address if host.is_lighthouse]
}
```

Run `TF_DF_API_KEY=<api-key> terraform plan` and you should see your host data!

```
$ TF_DF_API_KEY=dnkey-abc123 terraform plan
╷
│ Warning: Provider development overrides are in effect
│
│ The following provider development overrides are set in the CLI configuration:
│  - rfletcher/definednet in /path/to/go/bin
│
│ The behavior may therefore not match any released version of the provider and applying changes may cause the state to become incompatible with published releases.
╵
data.definednet_hosts.all: Reading...
data.definednet_hosts.all: Read complete after 0s

Changes to Outputs:
  + first_host     = {
      + created_at      = "2024-05-27T18:05:44Z"
      + id              = "host-ABC123"
      + ip_address      = "100.100.0.11"
      + is_blocked      = false
      + is_lighthouse   = false
      + is_relay        = false
      + listen_port     = 4242
      + name            = "rpi-1"
      + network_id      = "network-ABC123"
      + organization_id = "org-ABC123"
      + role_id         = "role-ABC123"
    }
  + lighthouse_ips = [
      + "100.100.0.1",
    ]
```
