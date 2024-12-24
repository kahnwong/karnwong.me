+++
title = "Bare metal works, until it doesn't. Hello, cloud."
date = "2023-03-24"
path = "/posts/2023/03/bare-metal-works-until-it-doesnt-hello-cloud"

[taxonomies]
categories = [ "infrastructure",]
tags = [ "aws", "gcp", "cloud"]

+++

<!-- # challenge you dealt with as a developer but later overcame it by using AWS -->

## Background

Ever wonder how websites (and everything in between) work? Chances are you can create a project running on your local machine. It works as you expected, but to let other people access it, you have to "deploy" it. For many years, to support a lot of request volumes you need to run your applications in a data center. These days this setup is known as on-premise.

## Architecture

Let's take a look at a simple e-commerce website architecture. This would involve frontend, backend, database. But if you have a feature where you allow users to batch update catalogs, you have to add redis / pub-sub as not to overload your backend.

This continues to work, except your website grew so popular that your cpu and memory usage is always at max capacity. You tried vertical scaling, but it's gotten to the point where it's no longer feasible for your budget. Oops. But wait, you can use a load balancer to spread requests throughout the nodes. This would work nicely, but it takes a lot of time to set it up properly. But you figured it out, so yay.

And you just happened to secured a major deal for your website, and they want a few customized features. You crank out a lot of features, but because it takes quite a while to deploy (manually) you made a lot of mistakes along the way, and you wish you can go back to release-once-a-month cycle.

## Cloud enters the chat

All these challenges are well known among on-premise users, and there are ways to overcome these. But using cloud reduce ops overhead and allows for faster time-to-market.

Let's break them down one-by-one to see how cloud can help:

### 1. Frontend, backend, database, etc

If you have to develop against local resources, you would need to spin up each component separately, but this means more context to management. Using cloud means you can develop against some resources that were already deployed on cloud.

For example, you can iterate on a UI change against a development instance of backend and database, so you can focus on developing the UI instead of worrying about setting up other components properly.

### 2. Scaling

It is a well-known fact with on-prem setup that you can't scale beyond your physical capacity, this essentially means if one fateful day, your marketing campaign works really well that a lot of people use your website, your traffic would spike up, in turn increasing resources consumption. If you're on-prem, your options are limited to either waiting it out or disable your website temporarily. But this would be catastrophic for your website's reputation.

If you use cloud, there are many ways to set up (auto) scaling, namely using container-based runtime and link running instances to a load balancer to spread out traffic. You can also temporarily increase your compute's provisioned resources during peak hours, then scale down later for cost optimization.

On AWS, you can use ECS for compute and define a load balancer. On GCP, Cloud Run can achieve the same thing.

### 3. CI/CD

From my experience and a lot of people I've talked to, deploying with on-prem setup usually means "copy deploy artifacts over to the server (manually)." Obviously this process is prone to human errors. You can set up Jenkins to achieve CI/CD, but this requires more overhead than using a cloud-managed build solutions.

With AWS, you can sync a repo with AWS Code Build - for building deploy artifacts, and deploy it via Code Pipeline. On GCP, similar services also available as well. Additionally, you can also do all this with your VCS providers (GitHub, Gitlab, etc) if you prefer it. Personally I would do CI/CD with via VCS providers. It requires more ops overhead, but yields better observability.

## Closing

A lot of people ask me which cloud to use. My answer would be: "use what other people around you use, so you can ask them when you have issues." Some clouds do certain things better than other clouds, that's up for you to evaluate. But if you're just starting out and don't have people who you can ask for support, AWS would be a better choice since it has significantly more users than other cloud providers. GCP is also good, but there's not as many users, which means unless you can understand official documentations, you're pretty much on your own.

Also, I did not mention Kubernetes. This is because unless you have a lot of services to manage or an ops team to support, Kubernetes' ops overhead would be too much for your scale.

And you can definitely achieve what you can do with cloud on on-premise system. But do you have enough people to setup and maintain it?
