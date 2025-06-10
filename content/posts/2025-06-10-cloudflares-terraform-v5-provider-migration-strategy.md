+++
title = "Cloudflare's terraform v5 provider migration strategy"
date = '2025-06-10'
path = '/posts/2025/06/cloudflares-terraform-v5-provider-migration-strategy'

[taxonomies]
categories = ['devops', 'migration', ]
tags = ['cloudflare', 'terraform']
+++

Back in February 2025, [Cloudflare announced that its terraform v5 provider is GA](https://developers.cloudflare.com/changelog/2025-02-03-terraform-v5-provider/). However, this release contains a lot of breaking changes, but I understand why it had to be this way - because it's less work if you generate a terraform provider via OpenAPI specs.

The [v4 to v5 provider upgrade guide is provided](https://registry.terraform.io/providers/cloudflare/cloudflare/latest/docs/guides/version-5-upgrade). However, it's not for the faint of heart, and it contains a lot of hacks. Moreover, some resources are renamed, which makes it even more tricky to perform a provider upgrade in-place, seeing it would fail the tfstate validation (trust me I've tried).

For a very small project, after a lot of trial and error I finally figured out that you have to remove all resources where the resource name changes in v5, then re-import them later.

However, I have a large terraform project, and this solution doesn't scale well. The good thing is that terraform state is in JSON format, so a little python-fu to extract resource names and ids should work, then initialize a new terraform project and slowly re-import everything back. You can copy the existing project, comment out all resource blocks and backend definition, then re-import cloudflare resources in chunks, and `terraform plan` periodically for sanity check.

## Python Utils Script

This is a little hacky, but it's a one-time kinda thing. You can modify this for other providers as well, but you'll have to adjust the blocks for resource name and ids extraction.

```python
import json

with open("terraform.tfstate") as f:
    state = json.load(f)

for resource in state["resources"]:
    resource_name = ""
    resource_id = ""

    if resource["mode"] == "managed":
        is_module = resource.get("module")
        if is_module:
            resource_name = (
                resource.get("module") + "." + resource["type"] + "." + resource["name"]
            )
        else:
            resource_name = resource["type"] + "." + resource["name"]

    resource_type = resource["type"]
    for instance in resource["instances"]:
        if resource_type == "cloudflare_record":
            instance_id = instance["attributes"]["id"]
            instance_zone_id = instance["attributes"].get("zone_id")
            resource_id = f"{instance_zone_id}/{instance_id}"
        elif resource_type == "cloudflare_api_token":
            resource_id = instance["attributes"]["id"]
        elif resource_type == "cloudflare_page_rule":
            instance_id = instance["attributes"]["id"]
            instance_zone_id = instance["attributes"].get("zone_id")
            resource_id = f"{instance_zone_id}/{instance_id}"
        elif resource_type == "cloudflare_r2_bucket":
            account = instance["attributes"]["account_id"]
            id = instance["attributes"].get("id")
            location = instance["attributes"].get("location")
            resource_id = f"{account}/{id}/{location}"
        elif resource_type == "cloudflare_pages_domain":
            account = instance["attributes"]["account_id"]
            project_name = instance["attributes"].get("project_name")
            domain = instance["attributes"].get("domain")
            resource_id = f"{account}/{project_name}/{domain}"
        elif resource_type == "cloudflare_pages_project":
            account = instance["attributes"]["account_id"]
            project_name = instance["attributes"].get("id")
            resource_id = f"{account}/{project_name}"
        if instance.get("index_key"):
            resource_name_loop = resource_name + '["' + instance["index_key"] + '"]'
            print(f"tf import '{resource_name_loop}' {resource_id}")
        else:
            print(f"tf import '{resource_name}' {resource_id}")
print("----")
```
