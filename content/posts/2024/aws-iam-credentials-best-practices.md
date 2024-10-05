---
title: AWS IAM credentials best practices
date: 2024-10-05T17:19:01+07:00
draft: false
ShowToc: false
images:
tags:
  - aws
  - devops
---

It's hard to escape AWS, seeing how prevalent it is in global internet infrastructure. Chances are, most websites you visit are hosted on AWS.

As a software engineer, you probably encounter AWS at certain point in your career, and while getting AWS IAM credentials to work locally during development (via `aws-cli`) would suffice, sometimes in production land, you might need some adjustments.

Having worked with AWS extensively, here are what I found to be useful to keep in mind while working with AWS IAM credentials.

## Disable `default` profile

When you first run `aws configure` it will ask you for `access key id` and `secret access key`. You can fill them, but you might accidentally delete or modify resources you didn't intend to. And rectifying the state can be time-consuming. Generally, I find it better to explicitly specify a specific AWS profile whenever I'm interacting with AWS. This is so there is an extra guardrails layer. Additionally, you probably need to use different AWS profiles for each project / service, might as well make it a habit.

## Define a separate `read-only` and `read-write` profile

Let's say I'm working on a service where users can upload files to S3. During development, I would need to use creds with `read-write` permissions, because I'm going to be doing a lot of tests, and until I finish the development, there would be a lot of manual uploads/deletes.

However, once things are deployed to production, I would spend more time investigating errors regarding this upload features, this would translate to "verify that $image exists on S3", and this only requires `read-only` permission.

While it's possible to use creds with `read-write` permissions, doing so on a prod bucket is like driving without wearing a seat belt. It works, but there is no extra layer of safety to prevent accidental oopsies, but it's not so fun recovering the said images.

## Grant permissions to groups, not users

This is generic for all RBAC situations, because it's not fun to synchronize permissions for all users in a group.

## Avoid supplying AWS IAM credentials in production

This applies to other cloud providers as well - generally cloud computes has an IAM role / service account attached, which means you don't have to explicitly inject IAM credentials to make it able to authenticate to AWS. And your security department would rejoice, because they don't have to deal with yet another "quarterly secrets rotation."
