+++
title = "Things to watch out for GCP SSL with Cloudflare DNS"
date = "2023-12-18"
path = "/posts/2023/12/things-to-watch-out-for-gcp-ssl-with-cloudflare-dns"

[taxonomies]
categories = ["infrastructure",]
tags = [  "gcp", "cloudflare", "ssl",]

+++

For our production workload, we deploy the workloads on Kubernetes, in which an ingress resource is created per each deployment. Resources in ingress are GCP Load Balancer and SSL Certificate. As for DNS, we use Cloudflare since it enables CDN without extra configurations on our part.

A few months after the deployment went live initially, we were informed that the website couldn't be accessed. Turns out GCP couldn't renew the SSL Certificate (error `FAILED_NOT_VISIBLE`.) Looking at GCP docs, turns out if the DNS couldn't be resolved to the Load Balancer IP, it couldn't provision/renew a certificate.

The fix? Disable proxy on Cloudflare records and wait until GCP successfully renews the certs.
