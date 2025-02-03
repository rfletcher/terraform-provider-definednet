# terraform-provider-definednet

An unofficial [Terraform](https://www.terraform.io) provider for the
[Defined Networking API](https://docs.defined.net/api/defined-networking-api/).

> [!WARNING]
> This is functional, but very incomplete! See [TODO](#TODO).

# Quick Start

This provider is only available locally, through Terraform's
[development overrides](https://developer.hashicorp.com/terraform/cli/config/config-file#development-overrides-for-provider-developers):

1. Install [Terraform](https://developer.hashicorp.com/terraform/install) and
   [Go](https://go.dev/doc/install)
2. Clone this repository
3. Make the plugin available to Terraform by editing `$HOME/.terraformrc` to
   include this configuration (be sure to use your actual *absolute* `$GOBIN`
   path -- often `$HOME/go/bin`):

        provider_installation {
          dev_overrides {
            "registry.terraform.io/rfletcher/definednet" = "/path/to/your/bin"
          }
          direct {}
        }

4. Run `make` to build and install
5. Run `TF_DN_API_KEY=<api-key> terraform -chdir=examples/quick-start plan`, and
   you should see some host data, read from the Defined API

```
╷
│ Warning: Provider development overrides are in effect
│
│ The following provider development overrides are set in the CLI configuration:
│  - rfletcher/definednet in /path/to/your/bin
│
│ The behavior may therefore not match any released version of the provider
│ and applying changes may cause the state to become incompatible with
│ published releases.
╵
data.definednet_hosts.all: Reading...
data.definednet_hosts.all: Read complete after 0s

Changes to Outputs:
  + sample_host      = {
      + created_at      = "2024-05-27T18:05:44Z"
      + id              = "host-ABC123"
      + ip_address      = "100.100.0.2"
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
  + node_name_ip_map = {
      + lighthouse-1  = "100.100.0.1"
      + rpi-1         = "100.100.0.2"
      + rpi-2         = "100.100.0.3"
    }
```

# TODO

- [x] Basic API client
- [x] API client <-> Terraform Provider plumbing
- [x] Initial data source (Hosts)
- [ ] Tests!
- [ ] Support for API pagination
- [ ] Support for API filtering
- [ ] Read support for more than just Hosts
- [ ] Maybe, one day, support for writing
