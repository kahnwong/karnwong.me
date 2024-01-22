---
title: Reduce operational costs with terraform
date: 2023-11-04T18:52:45+07:00
draft: false
ShowToc: false
images:
tags:
  - terraform
  - devops
  - finops
---

## Background

Think of websites you visit each day. Most likely they are hosted on a cloud provider such as AWS, GCP, Azure. The good news is it's very easy to create a simple deployment with a virtual machine, but for scalable and high-availability workloads, usual recommendations is to use a container-based runtime such as AWS ECS/EKS, GCP Cloud Run/GKE. These services also require more configurations than a simple VM deployment.

## The problem

Usually DevOps team would be provisioning required cloud resources through a cloud provider's web console (known as ClickOps), and if you're lucky, they write down the deployment details in a shared company documentation site, which has its own cost, especially if you require SSO.

This also means that, for each deployment pattern, there needs to be a corresponding entry in the documentation. Even more, some deployments require setting up environment variables to access a database or services requiring authentication, in which a DevOps engineer would also have to specify in the documentation where the secrets/configs are stored, since environment variables are dynamic, and they are tied to a deployment environment per each service.

This also poses another challenge, that a deployment steps outlined in the documentation can be out-of-date, since it requires an engineer to update the steps manually. If this keeps going on, no one would trust the documentation, which can result in a massive operation overhead (from tracking down the actual deployment steps, locate the secrets, etc) and knowledge loss if engineers leave the team.

## The operational costs

Assuming a team of DevOps consists of:

- 8 engineers
- 3 deployments per engineer per day
- $140K salary  -> 1 man-minute costs `140,000 USD / 12 months / 20 days / 8 hours / 60 minutes` = `1.2 USD/minute`

Provisioning cloud resources via a provider's web console has following costs:

- Setting up a shared documentation website; Confluence costs `6.05 USD * 8 engineers * 1 month` = `48.4 USD / month`
- Tracking down deployment steps, locating secrets, making sure that environment variables / secrets are correct: `3 services * 30 minutes * 8 engineers * 1.2 USD/minute * 20 days` = `17,280 USD / month`
- Updating deployment steps in the documentation: `3 services * 30 minutes * 8 engineers * 1.2 USD/minute * 20 days` = `17,280 USD / month`

## What about other costs?

Sometimes the finance department would inquire the cost of each service, in which a DevOps engineer would have to comb through the documentation, list cloud services used, then calculate the cost. This process can take around 15-30 minutes, and since this is done manually, some errors might occur due to human errors when obtaining resource types or looking up the costs.

Assuming there are 100 services, costs are calculated each quarter, and it takes an engineer 15 minutes to calculate the cost of each service: `100 services * 15 minutes * 1 engineers * 1.2 USD/minute * 4 quarters` = `7,200 USD / year`

## The solution

[Terraform](https://www.terraform.io/) is an IaC (infrastructure as code) tool, which means you can declare cloud resources programmatically, and you can uses it to apply cloud configurations. It can also act as a living documentation, since terraform code directly translates to desired cloud resources. This would eliminate the need for a documentation instance, tracking down steps/configurations and updating the documentation manually. Plus, it takes almost no time to perform a cost breakdown by using [Infracost](https://www.infracost.io/) CLI against a terraform project.

In total, you would save `48.4 USD + 17,280 USD + 17,280 USD` = `34,608.4 USD / month`.

In a single year, this means `34,608.4 USD / month * 12 months` = `415,300.8 USD / year` can be saved!

Very crazy!
