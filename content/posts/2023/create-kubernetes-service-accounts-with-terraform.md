---
title: Create Kubernetes service accounts with Terraform
date: 2023-08-01T19:49:07+07:00
draft: false
ShowToc: false
images:
tags:
  - devops
  - kubernetes
  - terraform
  - security
---

Sometimes you'll have to grant other people (or entities) access to your Kubernetes cluster. Easiest is you can give them your admin credentials, but this is similar to giving your house key to a friend, when they only need access to your living room.

You can give them different keys, depending on access level required. Those could be `readonly` access to view services status, or `deploy` service account that can create/update services.

You can mix and match these terraform configs, but most importantly, you should read up on Kubernetes RBAC before diving in further. You've been warned!

## Deploy service account

```hcl
# ------------------------ service account ------------------------ #
# https://github.com/hashicorp/terraform-provider-kubernetes/issues/1943#issuecomment-1369546028
resource "kubernetes_secret" "sa_github" {
  metadata {
    annotations = {
      "kubernetes.io/service-account.name" = kubernetes_service_account.sa_github.metadata.0.name
    }
    namespace = "default"
    name      = "${kubernetes_service_account.sa_github.metadata.0.name}-token"
  }
  type                           = "kubernetes.io/service-account-token"
  wait_for_service_account_token = true
}
resource "kubernetes_service_account" "sa_github" {
  metadata {
    name      = "sa-github"
    namespace = "default"
  }
}

# ------------------------ cluster role ------------------------ #
resource "kubernetes_cluster_role" "allow_deploy" {
  metadata {
    name = "cr-allow-deploy"
  }

  # https://stackoverflow.com/a/70669341
  rule {
    api_groups = [""]
    resources  = ["pods"]
    verbs      = ["list", "get", "watch", "create", "delete"]
  }
  rule {
    api_groups = [""]
    resources  = ["pods/exec"]
    verbs      = ["create"]
  }
  rule {
    api_groups = [""]
    resources  = ["pods/log"]
    verbs      = ["get"]
  }
  rule {
    api_groups = [""]
    resources  = ["pods/attach"]
    verbs      = ["list", "get", "create", "delete", "update"]
  }
  rule {
    api_groups = [""]
    resources  = ["secrets"]
    verbs      = ["list", "get", "create", "delete", "update"]
  }
  rule {
    api_groups = [""]
    resources  = ["configmaps"]
    verbs      = ["list", "get", "create", "delete", "update"]
  }
  rule {
    api_groups = [""]
    resources  = ["services"]
    verbs      = ["list", "get", "create", "delete", "update", "patch"]
  }
  rule {
    api_groups = ["apps"]
    resources  = ["deployments"]
    verbs      = ["list", "get", "create", "delete", "update", "patch"]
  }
}

# ------------------------ cluster role binding ------------------------ #
resource "kubernetes_cluster_role_binding" "allow_deploy" {
  # https://github.com/kubernetes/kubernetes/issues/30924#issuecomment-240887810
  metadata {
    name = "crb-allow-deploy-default"
  }
  role_ref {
    api_group = "rbac.authorization.k8s.io"
    kind      = "ClusterRole"
    name      = kubernetes_cluster_role.allow_deploy.metadata.0.name
  }
  subject {
    kind      = "ServiceAccount"
    name      = kubernetes_service_account.sa_github.metadata.0.name
    namespace = "default"
  }
}
```

## Readonly service account

```hcl
# ------------------------ service account ------------------------ #
# https://github.com/hashicorp/terraform-provider-kubernetes/issues/1943#issuecomment-1369546028
resource "kubernetes_secret" "foo" {
  metadata {
    annotations = {
      "kubernetes.io/service-account.name" = kubernetes_service_account.foo.metadata.0.name
    }
    namespace = "default"
    name      = "${kubernetes_service_account.foo.metadata.0.name}-token"
  }
  type                           = "kubernetes.io/service-account-token"
  wait_for_service_account_token = true
}
resource "kubernetes_service_account" "foo" {
  metadata {
    name      = "sa-foo"
    namespace = "default"
  }
}

# ------------------------ cluster role ------------------------ #
resource "kubernetes_cluster_role" "readonly" {
  metadata {
    name = "cr-readonly"
  }

  rule {
    api_groups = [""]
    resources  = ["namespaces", "pods", "pods/log"]
    verbs      = ["get", "list", "watch"]
  }
}

# ------------------------ cluster role binding ------------------------ #
resource "kubernetes_cluster_role_binding" "readonly" {
  metadata {
    name = "crb-readonly"
  }
  role_ref {
    api_group = "rbac.authorization.k8s.io"
    kind      = "ClusterRole"
    name      = kubernetes_cluster_role.readonly.metadata.0.name
  }
  subject {
    kind      = "ServiceAccount"
    name      = kubernetes_service_account.foo.metadata.0.name
    namespace = "default"
  }
}
```
