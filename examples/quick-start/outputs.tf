# a sample host, to show the complete set of fields available
output "sample_host" {
  value = data.definednet_hosts.all.hosts[0]
}

# an example of filtering results on the client side
output "lighthouse_ips" {
  value = [
    for host in data.definednet_hosts.all.hosts :
      host.ip_address
      if host.is_lighthouse
  ]
}

# a simple data transformation
output "node_name_ip_map" {
  value = {
    for host in data.definednet_hosts.all.hosts :
      host.name => host.ip_address
  }
}
