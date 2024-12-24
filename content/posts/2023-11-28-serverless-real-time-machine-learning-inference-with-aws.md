+++
title = "Serverless real-time machine learning inference with AWS"
date = "2023-11-28"
path = "/posts/2023/11/serverless-real-time-machine-learning-inference-with-aws"

[taxonomies]
categories = [ "mlops",]
tags = [ "aws", "machine-learning",]

+++

For a machine learning project, usually it is divided into two main categories: research and production. For research ML project, the model would be created and used locally on a researcher's machine. For a production ML project, a deployment would be involved. Usual pattern is to create a service to load a model, accept input, then return a prediction.

Production ML is also divided into two main patterns: batch or real-time. For batch inference, a job would be triggered on an interval to pre-calculate predictions, then store somewhere. As for real-time inference, it is more tricky, since this involves web application architecture (at least the data and application tier).

Let's assume there is a website where you can identify animals by uploading a photo. The model can expect a mandatory coordinates point to narrow down the possible list of animals. Ideally there would also be a database containing a list of animals in given geographical area. This service isn't used all the time, but during certain periods there are traffic spikes.

Together, during an inference it would look like this:

1. ML service reads input image
2. Extract coordinates from input image
3. Obtain relevant features from database
4. Obtain prediction from model
5. Return prediction

There are many ways to set up this ML inference architecture, but for less operational overheads, serverless offerings can be utilized. On AWS the architecture would look like this:

- AWS Lambda for ML service, since it automatically scales based on amount of incoming requests.
- AWS Aurora Serverless for storing data used during feature engineering step.
- AWS ElastiCache Serverless for caching features, to reduce lookup time during feature engineering.

Given the nature of sporadic service usage, but might have traffic spikes, this setup allows scaling, and you don't have to pay for compute when services are idle during off-peak hours.

To elaborate, when there are a lot of concurrent requests, Lambda would scale on its own. But the bottleneck would be the database, since it doesn't always scale orthogonally to Lambda. However, every time an ML service fetches data from database, it doesn't necessarily looking for new data. In most cases, it would look up the same data because users tend to submit images taken from certain geographical areas, which can be their hometown, current location, or tourist spots. An in-memory cache such as Redis (in AWS it's ElastiCache) can be used to reduce database loads by caching frequently-accessed data. Previously the full serverless setup isn't possible because AWS didn't offer serverless in-memory database, but [now it does](https://aws.amazon.com/blogs/aws/amazon-elasticache-serverless-for-redis-and-memcached-now-generally-available/)!

You can also combine this with AWS ECS on Fargate for an actual application tier, which means you can fully focus on your application and not the operational overhead. And storage can be S3, in which AWS manages everything for you. Print logs here and there and send it to CloudWatch, and now you have observability (ECS and Lambda automatically log metrics!)

The possbilities!
