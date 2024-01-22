---
title: Cost optimization with kubernetes
date: 2023-04-01T16:34:11+07:00
draft: false
ShowToc: false
images:
tags:
  - devops
  - kubernetes
  - prometheus
  - opencost
  - finops

---

Correction 2023-07-02: fix homelab specs and corresponding AWS EC2 instance class (it's actually 32GB RAM, not 64GB)

Congratulations, you managed to successfully deployed a few services on kubernetes! But this is not the end 👀. Unfortunately money doesn't grow on trees, and if you can't justify your infra expenses, finance department won't be happy.

If you're using Terraform, you can use [Infracost](https://www.infracost.io/) to create a cost report. Pretty nifty. But what about kubernetes? Given cost reporting is a basic feature, kubernetes is no exception.

Enter [OpenCost](https://www.opencost.io/), a vendor-neutral open source project for measuring and allocating infrastructure and container costs in real time. This also means it is vendor-agnostic. Meaning as long as it's kubernetes, it would work regardless of which cloud it's on (and it works with on-prem setup as well 😉).

Under the hood, it utilizes [prometheus metrics], in conjunction with [provisioned resources for containers - defined in deployment manifest], to calculate cost based on [cloud pricing per compute unit - you can adjust this later].

You can follow OpenCost install instructions [here](https://www.opencost.io/docs/install). Wait for a few days for prometheus to collect usage metrics, and check out OpenCost dashboard. This is what mine looks like:

![picture 1](images/4d695173c90ef2db997019774a93863847a39368e35bfdaf48d79b76acca515b.webp)

Notice `efficiency` column, this tells how well your resources are being utilized. If you see a low number here, consider using function-as-a-service or other cloud compute where pricing is calculated per usage, not active time (something like gcp's cloud run).

## Boring math part 😴

So that's `$35.57 / 6 days`, which would be around `$5.9 / day`. Per year this would cost `$5.9 * 52 = $306.8 / year` 😱.

My current home server setup is around `$880`, so that's around `3 years` before it would break even 🤣.

But what if we're talking about raw compute price? My server is `8 cores, 32GB RAM, 1TB SSD`, cheapest compute on AWS with similar specs per year is `t3a.2xlarge`, which is `$3307 / year`.

Pricing is hard...
