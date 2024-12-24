+++
title = "Self-hosting primer"
date = "2021-08-22"
path = "/posts/2021/08/self-hosting-primer"

[taxonomies]
categories = ["homelab",]
tags = [ ]

+++

Self-hosting is a practice for running and managing websites / services using your own server. Some people do this because they are concerned about their privacy, or some services are free if they host it themselves. Below are instructions for how to do self-hosting (also applies to hosting your own website too).

## Requirements

- Domain name
- Server (can be your own computer at home or VPS)

## Instructions

1. Set up and secure the server (set up password, disable password login (which means you can only login via SSH key), etc.)
2. Deploy a website on your server (follow instructions for each service. I recommend deploy via Docker).
3. If you are using a server at home which has dynamic IP, setup DDNS (I recommend duckdns.org, since it has very fast TTL).
4. Go to your domain name registrar, under DNS, add a CNAME record for your desired subdomain, and set the value to your duckdns.org domain.
5. On your server, install a webserver for reverse-proxy. I recommend nginx or Caddy.
6. Create a virtual host config for your website in your webserver of choice.
7. On your router configuration page, under port forwarding, create two entries for port 80 and 443.
8. Wait for a few minutes for the DNS to be updated, and you should be able to access your website from the specified domain.

As for actual implementation, I suggest you read a few articles for each step, so you can get the overall idea of what's to be done. Generally, the common steps should be the same across all articles, since that's the "baseline" for each process.
