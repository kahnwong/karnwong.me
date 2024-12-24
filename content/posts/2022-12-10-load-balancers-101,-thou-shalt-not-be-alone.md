+++
title = "Load balancer 101, thou shalt not be alone"
date = "2022-12-10"
path = "/posts/2022/12/load-balancers-101,-thou-shalt-not-be-alone"

[taxonomies]
categories = [ "infrastructure",]
tags = []

+++

Scaling, the dreaded word among developers, because this means more complexity. But why do we need scaling?

Imagine a super busy corner store. During early mornings, there might not be a lot of customers, so one cashier might be enough to handle all customers. But during afternoons or evenings, more customers would flock to the store, and our only cashier couldn't checkout fast enough, and this means losing potential customers.

Is there a way to solve this? Good news is that's a "yes." However, there are a few implications. You could replace the only cashier with a cashier who can do a checkout faster. But, you could also hire more cashiers as temps during peak hours, and this means cheaper cost per customer compared to maxing out the only cashier. Maximizing a single resource is "vertical scaling", whereas adding more resources of the same caliber is called "horizontal scaling." Couldn't find a more apt name myself.

I did mention about complexity. Vertical scaling is less complex than horizontal scaling, because you don't need to coordinate the resources. Because if you have three instances of a website running, but there's only one entrypoint (the website's domain name) then how are you going to distribute the requests to each instance?

Load balancer to the rescue. Essentially it acts as a reception front, where it sends incoming requests to available instances. Think of a busy check-in queue in a hotel. You would be waiting in a single line, and when a counter frees up the next person would go there.

So how do we actually do this with a website? Common wisdom says using cloud, add a task to container runtime engine, set scaling policy, and route them to a load balancer.

There's a [load balancer challenge from pulumi](https://www.pulumi.com/challenge/holiday-shopping/) you could try out, since their example is tested to be working, which would save a lot of time from stitching up stuff from online articles and official documentations. Note that in this challenge, task scaling isn't implemented, so get your itchy fingers working and implement it!
