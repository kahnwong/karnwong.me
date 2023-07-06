---
title: Secrets management with SOPS, AWS Secrets Manager and Terraform
date: 2021-11-30T20:11:12+07:00
draft: false
ShowToc: true
images:
tags:
  - devops
  - aws
  - github
  - terraform
  - recommended
---

Correction 2023-07-06: I only recently realized SSM and Secrets Manager are not the same.

At my organization we use sops to check in encrypted secrets into git repos. This solves plaintext credentials in version control. However, say, you have 5 repos using the same database credentials, rotating secrets means you have to go into each repo and update the SOPS credentials manually.

Also worth nothing that, for GitHub actions, authenticating AWS means you have to add repo secrets. This means for all the repos you have CI enabled, you have to populate the repo secrets with AWS credentials. When time comes for rotating the creds, you'll encounter the same situation as above.

I did some research and consensus for AWS / Terraform setup is to: encrypt secrets via SOPS, and use Terraform to create AWS secret entries. That way, you have a trail for credentials. This setup means:

1. You don't have to populate repos with AWS creds, instead supplying an ARN role instead.
2. You don't have to change credentials in projects, since they all get the secrets from AWS Secrets Manager.

## Implementation

Repo here: <https://github.com/kahnwong/terraform-sops-ssm>

### 1. Bootstrap Terraform

```hcl
terraform {
  required_providers {
    sops = {
      source  = "carlpett/sops"
      version = "0.6.3"
    }
  }
}

provider "aws" {
  region  = "ap-southeast-1"
  profile = "playground"
}

provider "sops" {}

terraform {
  required_version = ">= 1.0"
}
```

### 2. Create KMS key for SOPS

<https://github.com/mozilla/sops/#kms-aws-profiles>

```hcl
resource "aws_kms_key" "sops" {
  description = "Keys to decrypt SOPS encrypted values"
}
resource "aws_kms_alias" "sops" {
  name          = "alias/sops"
  target_key_id = aws_kms_key.sops.key_id
}
```

### 3. Create secrets

Create a folder named `secrets`, inside it create JSON files and encrypt each with sops.

```hcl
locals {
  secrets = toset([
    "db-foo",
  ])
}

data "sops_file" "sops_secrets" {
  for_each    = local.secrets
  source_file = "secrets/${each.key}.sops.json"
}
# aws keeps the secrets for 7 days before actual deletion. consider using random names during test
resource "aws_secretsmanager_secret" "ssm_secrets" {
  for_each = local.secrets
  name     = each.key
}
resource "aws_secretsmanager_secret_version" "ssm_secrets" {
  for_each      = local.secrets
  secret_id     = aws_secretsmanager_secret.ssm_secrets["${each.key}"].id
  secret_string = jsonencode(data.sops_file.sops_secrets["${each.key}"].data)
}
```

### 4. Create IAM policy for SSM access

```hcl
data "aws_iam_policy_document" "secrets_ro" {
  statement {
    actions = [
      "secretsmanager:GetResourcePolicy",
      "secretsmanager:GetSecretValue",
      "secretsmanager:DescribeSecret",
      "secretsmanager:ListSecretVersionIds"
    ]
    resources = [
      "arn:aws:secretsmanager:ap-southeast-1:$AWS_ACCOUNT_ID:secret:*",
    ]
  }
  statement {
    actions = [
      "secretsmanager:ListSecrets"
    ]
    resources = ["*"]
  }
}
resource "aws_iam_policy" "secrets_ro" {
  name   = "secrets_ro"
  path   = "/"
  policy = data.aws_iam_policy_document.secrets_ro.json
}
```

### 5. Create IAM user for local dev

You shouldn't supply AWS credentials for deployment, since you can grant access via IAM roles instead.

```hcl
resource "aws_iam_user" "playground-prod-dev" {
  name = "playground-prod-dev"
  path = "/users/"
}
resource "aws_iam_access_key" "playground-prod-dev" {
  user = aws_iam_user.playground-prod-dev.name
}
```

#### Grant IAM user access to Secrets

```hcl
resource "aws_iam_user_policy_attachment" "playground-prod-dev" {
  user = aws_iam_user.playground-prod-dev.name

  for_each = toset([
    aws_iam_policy.secrets_ro.arn,
  ])
  policy_arn = each.value
}
```

### 6. Create IAM role for Lambda

```hcl
resource "aws_iam_role" "lambda_role" {
  name = "lambda_role"
  path = "/sa/"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        "Effect" : "Allow",
        "Principal" : {
          "Service" : "lambda.amazonaws.com"
        },
        "Action" : "sts:AssumeRole"
      },
    ]
  })

  managed_policy_arns = [
    aws_iam_policy.secrets_ro.arn,
  ]

  inline_policy {
    name = "create_cloudwatch_logs"
    policy = jsonencode({
      "Version" : "2012-10-17",
      "Statement" : [
        {
          "Action" : [
            "logs:CreateLogGroup",
            "logs:CreateLogStream",
            "logs:PutLogEvents"
          ],
          "Effect" : "Allow",
          "Resource" : "*"
        },
      ]
    })
  }
}
```

### 7. Create IAM role for GitHub actions

Need to create OIDC so GitHub can assume AWS roles

```hcl
resource "aws_iam_openid_connect_provider" "github" {
  url             = "https://token.actions.githubusercontent.com"
  client_id_list  = ["sts.amazonaws.com"]
  thumbprint_list = ["a031c46782e6e6c662c2c87c76da9aa62ccabd8e"]
}
```

Assume role policy is on per-repo basis

```hcl
locals {
  repositories = [
    "terraform-sops-ssm",
  ]
}
data "aws_iam_policy_document" "github_actions_assume_role" {
  statement {
    actions = ["sts:AssumeRoleWithWebIdentity"]
    principals {
      type        = "Federated"
      identifiers = [aws_iam_openid_connect_provider.github.arn]
    }
    condition {
      test     = "ForAnyValue:StringLike"
      variable = "token.actions.githubusercontent.com:sub"
      values   = [for v in local.repositories : "repo:kahnwong/${v}:*"]
    }
  }
}
```

Finally attach the above policy to role

```hcl
resource "aws_iam_role" "playground-prod-github" {
  name = "playground-prod-github"
  path = "/sa/"

  assume_role_policy = data.aws_iam_policy_document.github_actions_assume_role.json

  managed_policy_arns = [
    aws_iam_policy.secrets_ro.arn,
  ]
}
```

The end 🎉

## Bonus

- [GitHub actions with AWS login](https://github.com/kahnwong/terraform-sops-ssm/blob/master/.github/workflows/deploy.yml)
- [AWS Lambda with SSM template](https://github.com/kahnwong/terraform-sops-ssm/tree/master/app)
