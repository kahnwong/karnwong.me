---
title: IaC is cool, until someone specified the wrong machine size 💸
date: 2022-11-03T00:55:31+07:00
draft: false
ShowToc: false
images:
tags:
  - devops
---

Back in the day, there was no cloud. If you want a lot of computing power, you need to build your own data center, and this is very expensive. Then cloud happened, and suddenly you can work with a lot of flexibility like you couldn't before. Want to try out a small deployment? Sure! If your workload is heavier you can always increase the VM specs.

It would almost be the end of the story, until someone realize "oh my God how did I set up VPC peering with another VPC, and which route table did I choose? I need to replicate it with another client and I forgot how". THe issue isn't that working with the cloud's web ui doesn't work, rather it doesn't leave a trail for you, so over time the knowledge of how things were set up are lost, gone forever. I'm sure you can always mitigate this by documenting things, but then it can happen that someone modified the infrastructure _without_ updating the documentation. Can't blame them, it's only normal when you have to synchronize things manually.

So people came up with IaC - infrastructure as code. Basically this means, you write declarative blocks of what a resource's configuration should be, and how it's connected to another resources. Then you "apply" the configuration on your cloud provider. If you are researching or want to get something up and running fast, I wouldn't suggest you use IaC. But if you have to modify the configuration, you might find it easier to maintain in the long run, plus you can always use `git diff` to show the changes between each commit.

And as how most things work, this certainly is not the end, because what if you have to create a new VM fleet, a small cluster for a research project. IaC can do this, but what if you were working late at night and you copied the wrong instance size into your IaC config - what could be 50 USD/month small cluster could turn out to be a 2000 USD / month large cluster, and it won't even be utilized that much because your research workload is very light 💸💸💸💸💸.

Luckily humans excel at adaptation, so we came up with "IaC policy", where you can define a set of constraints that your resources should comply to. This could be anything from "only instance smaller than $x is allowed", or "total monthly cost most not exceed $y". It can also be "maximum storage allowed is 500GB", or even "daily backup is required for $z resource". The world is not that gloomy, a lot of people made sure of that 🚀.

And I know you are eager to get your hands on at policy as code thing, so head to [a pulumi challenge](https://www.pulumi.com/challenge/one-quickstart/) and take a stab at it!
