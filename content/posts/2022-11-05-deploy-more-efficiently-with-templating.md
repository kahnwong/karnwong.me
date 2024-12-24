+++
title = "Deploy more efficiently with templating"
date = "2022-11-05"
path = "/posts/2022/11/deploy-more-efficiently-with-templating"

[taxonomies]
categories = ["devops",]
tags = [ ]

+++

You are building a website, it's a simple frontend that needs to call the database for [total lead drops this week]. Your website is still at an infancy stage, with only a few features. At this point, you contemplate whether you need a proper backend or not. But to deploy a backend properly, it would involve docker, backend database, persistence storage, DNS, load balancer, among other things. But it looks like you don't have enough time, so you decide to go with serverless, since it takes less time to implement and you don't have to worry about scaling.

So you go about creating a serverless function to fetch the information from the database. But the frontend can't call the function directly, so you have to expose the function through an api gateway, so frontend can talk to it. You are happy with the current setup, since it relives you a lot of maintenance effort of creating a proper backend.

But your business is still trying to find its footing, so you had an idea of tweaking the query every so often. Then you are getting more annoyed at the fact that, every time you have to update the query, you have to go to the cloud console and manually update the query in your serverless app. You feel like there should be an easier way of just updating a few lines of code. You came upon a black magic called CI/CD. You tinkered with it a bit, then you smiled because this means you don't ever have to log in the cloud console ever again as long as you hook up the CI/CD to your serverless code properly.

A few weeks went on, you couldn't be happier. Then you had an idea for a few more features, but at this point you still are not ready for deploying a proper backend yet, so you have to create a few serverless functions for each feature. Suddenly you froze, having realized that you forgot how you setup the first serverless function, not to mention how to connect it to api gateway, and hook it up to CI/CD. Sweats tricking down your forehead, bracing for the eternal doom of doing everything by hand manually every time you update the code.

Luckily it's 2022, and the devops tooling space improved significantly compared to the dark ages of VMware. Apparently smart peoples utilize a certain form of template to spin up serverless functions, since it's been dawned on them that most of the configuration differences are your code, not the infrastructure setup.

So basically if you wish to create a serverless api to interface with via frontend, this magical template would:

1. Create a serverless function
2. Inject your code into the function
3. Hook up the function with api gateway
4. Also sprinkle a magic CI/CD hook so the whole flow would be triggered again from start to finish after a git push

Most leading IaC tools have this feature, but what if you are not familiar with HashiCorp configuration language? Then you might find such tools with normal programming languagae interface more familiar, in which Pulumi is a contender in this space. HashiCorp also recently released CDK to accomplish the same thing if you want to give it a go as well. But for a start, you could check out [Pulumi's Deployments Mini-Challenge](https://www.pulumi.com/challenge/deployments/) to try out deploying a serverless function with templating and CI/CD. For Terraform CDK, check out [this repo](https://github.com/cdktf/cdktf-integration-serverless-example).
