# Example: quick-start

A simple module which reads host details from the Defined Networking API, and
displays them in the Terraform plan.

1. [Build the `terraform-provider-definednet` binary](../../README.md#quick-start)
2. Run `terraform plan`, setting the equired variable, and it should show a plan
   similar to the one below

```
$ terraform plan -var definednet_api_key=<key>
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
