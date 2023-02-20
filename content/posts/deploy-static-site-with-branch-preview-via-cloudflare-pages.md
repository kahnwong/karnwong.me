---
title: Deploy static site with branch preview via Cloudflare Pages
date: 2022-10-05T02:23:06+07:00
draft: false
ShowToc: false
images:
tags:
  - cloudflare
  - devops
  - terraform
---

_Updated 2023-02-20_: update terraform code

For frontends, if no server-side rendering is required, we can deploy it as a static site. If you already use GitHub, you might be familiar with GitHub Pages. One common use case is to deploy your personal landing page / blog via GitHub Actions.

Interestingly enough, this might cause problems if you are working in a team. For example, if you are working on a UI change, and you need to have someone else approve the changes, they would need to build the site locally to do so.

Luckily, "branch preview" feature exists. Essentially it's a mechanism to generate a preview site for every pull request. We are going to use Cloudflare Pages for this (alternatives are Vercel, etc.)

You should have a github repo with the source code for your site. Then,

## Terraform

```hcl
resource "cloudflare_pages_project" "this" {
  account_id        = var.account_id
  name              = var.project_name
  production_branch = var.production_branch

  lifecycle {
    ignore_changes = [deployment_configs, source]
  }
}

resource "cloudflare_record" "this" {
  depends_on = [cloudflare_pages_project.this]

  name    = var.subdomain
  proxied = true
  ttl     = 1
  type    = "CNAME"
  value   = cloudflare_pages_project.this.subdomain
  zone_id = var.zone_id
}

resource "cloudflare_pages_domain" "this" {
  depends_on = [cloudflare_record.this]

  account_id   = var.account_id
  project_name = var.project_name
  domain       = var.domain_name
}
```

~~Notice that `cloudflare_pages_domain` is commented out, this is because it has a bug that would throw "error creating domain". Although behind the scenes it creates a subdomain successfully (and is actually working with cloudflare pages).~~

### GitHub Actions

We are going to use github actions as a runner to build and deploy the site.

```yaml
name: Deploy

on:
  push:
  workflow_dispatch:

concurrency:
  group: environment-${{ github.ref }}
  cancel-in-progress: true

jobs:
  publish:
    runs-on: ubuntu-latest
    permissions:
      contents: read
      deployments: write
    name: Publish to Cloudflare Pages
    steps:
      - name: Checkout
        uses: actions/checkout@v3
      # ---------- build ----------
      - name: Setup node
        uses: actions/setup-node@v3.5.0
        with:
          node-version: "18"
          cache: "yarn"
      - name: Install requirements
        run: yarn install
      - name: Build
        run: yarn run build
      # ---------- publish ----------
      - name: Publish to Cloudflare Pages
        uses: cloudflare/pages-action@1
        with:
          apiToken: ${{ secrets.CLOUDFLARE_API_TOKEN }}
          accountId: ${{ secrets.CLOUDFLARE_ACCOUNT_ID}}
          projectName: YOUR_PROJECT_NAME
          directory: $YOUR_BUILD_DIR
          gitHubToken: ${{ secrets.GITHUB_TOKEN }}
```

Notice the "build" section, you can adjust this per your site's setup.

This actions would work for both a push to master and a pull request

![cloudflare pages branch preview](/images/2022-10-05-02-40-24.png)
