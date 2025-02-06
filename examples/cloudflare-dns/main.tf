# fetch all hosts from Defined
data "definednet_hosts" "all" {}

# build a { name => ip } map from the results
locals {
  host_ip_map = {
    for host in data.definednet_hosts.all.hosts :
      host.name => host.ip_address
  }
}

# look up the target zone on Cloudflare
data "cloudflare_zone" "target" {
  name = var.cloudflare_zone_name
}

# create a <name>.nebula.example.com A record for each Defined host
resource "cloudflare_record" "nebula_host" {
  for_each = local.host_ip_map

  zone_id = data.cloudflare_zone.target.id

  name    = "${each.key}.nebula"
  content = each.value
  type    = "A"
  ttl     = 60 # set low for demonstration purposes
  proxied = false
}
