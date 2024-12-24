+++
title = "Cloudflare DNS management with Terraform"
date = "2022-11-20"
path = "/posts/2022/11/cloudflare-dns-management-with-terraform"

[taxonomies]
categories = ["infrastructure",]
tags = [  "cloudflare", "terraform",]

+++

I self hosted a lot of services, sometimes I try out a few apps that would get deleted within the same day. All this requires setting up CNAME for reverse-proxy (because I want to make sure there's no funny reverse-proxy shenanigans going on, for future reference).

I can always log into Cloudflare console and manually add CNAME entries, but this is getting too tiresome since all I really need is another CNAME with the same config as the rest of the CNAMEs - pointing to the same DNS for my homelab. Cue lightbulb moment when I realize I can use Terraform to set it up.

It's as simple as:

```terraform
locals {
  selfhosted_proxied = toset([
    "service_a", "service_b", "service_c",
  ])
  selfhosted_non_proxied = toset([])
}
locals {
  proxied_dict     = { for name in local.selfhosted_proxied : name => true }
  non_proxied_dict = { for name in local.selfhosted_non_proxied : name => false }
  selfhosted_dns   = merge(local.proxied_dict, local.non_proxied_dict)
}

resource "cloudflare_record" "selfhosted_dns" {
  for_each = local.selfhosted_dns
  name     = each.key
  proxied  = each.value
  ttl      = 1
  type     = "CNAME"
  value    = var.ddns
  zone_id  = var.zone_id
}
```

Told you it was really easy 😎.
