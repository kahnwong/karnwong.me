---
title: Reasons why you shouldn't use programming languages for IaC
date: 2024-08-05T17:38:13+07:00
draft: false
ShowToc: false
images:
tags:
  - devops
  - iac
  - terraform
---

When it comes to IaC (infrastructure as code), most people might have heard of HashiCorp's Terraform (it uses HCL as DSL. Interestingly enough, Terraform also has its own CDK to translate programming languages into HCL), Pulumi or AWS CDK. The latter two support programming languages as DSL.

Mostly there are two camps:
- People who swear by HCL and think you shouldn't use programming languages for IaC
- People who don't see why you need to pick up a new language in order to use IaC, so they prefer using a programming language they already are familiar with instead

Both camps are not wrong, they are both valid. However, I want to share my take on why you should use HCL for IaC.

## Pros
The obvious perk is that you don't need to pick up a new language, which means you can get started right away to enjoy the benefits of IaC.

## Cons

However, when you are looking for resources or templates, you might find it harder because there are many different implementations to achieve the same thing, due to the nature of programming languages where it allows for flexibility. To create a single VM inside a VPC, you can run into many different examples, depending on how people define the IaC.

Additionally, if you are using less popular languages (golang, for instance) you might have fewer references than people who use javascript/typescript or python.

As for IaC codebase maintenance, you would have to find people who:
1. Can use your programming language of choice for IaC
2. Understands devops and infrastructure

Which would significantly reduce your candidate pools, whereas if you use HCL (a DSL, which means it's less complex than a programming language, and by definition, has less learning curve), you would not be limited by the fact that your candidate has to know a specific programming language.

The same argument can be made that you still have to find people who know HCL, but being a DSL, it takes less time to get familiar with it. (Some notes on this: some people find it really hard to pick up a new language/DSL/whathaveyou. So it's understandable if they prefer to use a familiar programming language for IaC, and there's nothing wrong with that.)
