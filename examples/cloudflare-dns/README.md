# Example: cloudflare-dns

A module which creates DNS records in Cloudflare for each of your Defined
Networking hosts.

> [!NOTE]
> There's a terraform bug when using a
> "[dev override](https://developer.hashicorp.com/terraform/cli/config/config-file#development-overrides-for-provider-developers)"
> for a provider that isn't also in their registry. If `terraform init` gives
> you an error, try commenting-out all of the references to `definednet` in this
> example, `init` again, and then un-comment.)

1. [Build the `terraform-provider-definednet` binary](../../README.md#quick-start)
2. Run `terraform init`, to install the cloudflare plugin
3. Run `terraform plan`, setting the three required variables, and it should
   show a plan similar to the one below

```
$ terraform plan \
  -var cloudflare_zone_name=example.com \
  -var cloudflare_api_token=<token> \
  -var definednet_api_key=<key>
╷
│ Warning: Provider development overrides are in effect
│ 
│ The following provider development overrides are set in the CLI configuration:
│  - rfletcher/definednet in /path/to/your/bin
│ 
│ The behavior may therefore not match any released version of the provider and
│ applying changes may cause the state to become incompatible with published
│ releases.
╵
data.definednet_hosts.all: Reading...
data.cloudflare_zone.target: Reading...
data.definednet_hosts.all: Read complete after 1s
data.cloudflare_zone.target: Read complete after 1s [id=abc123]

Terraform used the selected providers to generate the following execution plan.
Resource actions are indicated with the following symbols:
  + create

Terraform will perform the following actions:

  # cloudflare_record.nebula_host["lighthouse-1"] will be created
  + resource "cloudflare_record" "nebula_host" {
      + allow_overwrite = false
      + content         = "100.100.0.1"
      + created_on      = (known after apply)
      + hostname        = (known after apply)
      + id              = (known after apply)
      + metadata        = (known after apply)
      + modified_on     = (known after apply)
      + name            = "lighthouse-1.nebula"
      + proxiable       = (known after apply)
      + proxied         = false
      + ttl             = 60
      + type            = "A"
      + value           = (known after apply)
      + zone_id         = "abc123"
    }

Plan: 1 to add, 0 to change, 0 to destroy.

────────────────────────────────────────────────────────────────────────────────

Note: You didn't use the -out option to save this plan, so Terraform can't
guarantee to take exactly these actions if you run "terraform apply" now.
``` 
