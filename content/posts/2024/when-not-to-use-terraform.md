---
title: When (not) to use Terraform
date: 2024-10-05T19:32:29+07:00
draft: false
ShowToc: false
images:
tags:
  - devops
  - terraform
---

If we are talking about IaC, Terraform would be on the list. It made IaC popular and help a lot of companies maintain infrastructure at scale. Especially when you have multiple sets of infrastructures to maintain, Terraform can help you reduce the setup time tremendously via using Terraform Modules. Think of this like a function / class in programming languages.

I've been cranking out a lot of Terraform, a lot of trials and errors along the way. Picking other people's brains by reading a lot of blog posts, trawl community forums to see how other people use Terraform and what are their challenges. Below are summarizations of what I've experienced, in addition to what the general sentiments are.

## Terraform is great for foundational resources

Think of foundational resources as the first layer you have to build before you can deploy your apps. This would include databases, VPCs, virtual machines, storage volumes. These are rarely changed, or if you want to, downtimes can occur.

Out of all of these, it's probably safest to only modify database configs and virtual machine specs, since it results in the least downtime.

## For app deployment, YMMV

Although Terraform can also be used to deploy apps, it is not a good idea to have [foundational resources] in the same workspace as [app deployment].

This is because after initial app deployment, subsequent deployments should be done via CI/CD, this essentially means [something else that's not Terraform is updating the app deployment manifest][^1]. This translates to [your Terraform state will constantly be out-of-sync].

This especially is a problem if your Terraform config uses `latest` for image tag. If your devs just rolled back a deployment to a previous SHA, if you apply this Terraform workspace, chances are you are going to override the deployment rollback (and this would make the devs confused as heck as to why their deployment is suddenly on the latest version after they just did a rollback).

But it's never a no, because I still find it practical to initialize an app deployment in Terraform, but it's in a separate workspace on a per-app basis. This means it won't mess with other infra resources outside of app deployments.

[^1]: Technically you can use Terraform in CI/CD for app deployment, but that's another can of worms.

## If you can, use it to set up / provision IAM _users_ access

I have a friend who works a lot with auditors, and he told me that most orgs fail auditing because they don't have a process for IAM access provisioning. This means you can't track who has access to what, or when the permissions had been granted to particular users/groups.

It is very tedious, and it should be in a separate Terraform workspace. But if you can achieve this, you would have a single source-of-truth for RBAC.
