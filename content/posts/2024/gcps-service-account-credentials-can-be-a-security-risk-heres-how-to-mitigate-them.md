---
title: GCP's service account credentials can be a security risk. Here's how to mitigate them.
date: 2024-07-14T14:01:57+07:00
draft: false
ShowToc: false
images:
tags:
  - gcp
  - security
  - devops
---

If you look online, many sources would tell you that you should use service account to authenticate for GCP services. While this is true, it's not for all the cases.

## For local development, you should use Application Default Credentials

Imagine working in a team, and you have to work with Cloud Run, so you request your infra team for a service account. This looks good, but then your teammates also have to work with this service. They happen to be in a hurry, so you share your service account to your teammates. Now this can be a problem, because now there are multiple users who have access to this service account. It would be very tricky to trawl through the audit logs and identify which developer interact with cloud run, because the system only sees a single identity.

This is why for a local development workflow, application default credentials should be used to authenticate to GCP. Essentially you perform a login action via the gcloud command line and set ADC, it would then leave a credential file in your local filesystem. Other Google/GCP SDKs can then pick this up and use it to authenticate to GCP services. This also adds a benefit that no credentials have to be generated, meaning there's no cleanup overhead.

## For cross-perimeter authentication, only required personnel should be able to access the credentials

But if we are talking a deployment, it's possible that you need to have $serviceA talks to $serviceB. In GCP, you can assign service account to the compute directly, as this method doesn't require generating a credentials.

But there's also cases where you have to access GCP services from outside of GCP. For example, using Fivetran to read data from BigQuery. Unfortunately this method requires creating a credential tied to a service account. This can become a risk if you don't limit permission for who can create / view credentials. Following the principle of least privileges can in turn mitigate unauthorized actors from accessing service account credentials.
