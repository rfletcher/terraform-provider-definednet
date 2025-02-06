# terraform-provider-definednet

An unofficial [Terraform](https://www.terraform.io) provider for the
[Defined Networking API](https://docs.defined.net/api/defined-networking-api/).

> [!WARNING]
> This is functional, but very incomplete! See [TODO](#TODO).

# Quick Start

This provider isn't available from the Terraform Registry. To use it locally,
build the binary and configure Terraform's
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
5. Try out the ["quick start" example](examples/quick-start/README.md)

# TODO

- [x] Basic API client
- [x] API client <-> Terraform Provider plumbing
- [x] Initial data source (Hosts)
- [ ] Tests!
- [ ] Support for API pagination
- [ ] Support for API filtering
- [ ] Read support for more than just Hosts
- [ ] Maybe, one day, support for writing
