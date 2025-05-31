+++
title = 'My code search setup throughout the years'
date = '2025-05-31'
path = '/posts/2025/05/my-code-search-setup-throughout-the-years'

[taxonomies]
categories = ['platform']
tags = []
+++

## Initial Setup

Around 2022-23, I was browsing `awesome self-hosted` and came across `Sourcegraph`, a popular code search tool. I don't plan to pay, and they allowed self-hosting, so it was perfect for me at the time. Then you still can't search code in GitHub (which we used for work), so being able to index ever-growing repos was very important. My productivity skyrocketed and our teams were happy because they don't have to hunt down that specific line in god-knows-which-repo-and-which-file.

This went on for two years, then the license changes happened and we can't use self-hosted Sourcegraph with more than a handful of private repos anymore.

## Livegrep to the Rescue

I then discovered `livegrep`, a very lightweight code search solution. Previously with Sourcegraph, to make things easier, they bundled all services into a single container image, which is a good thing because it's very convenient. But with livegrep, you need three components: an indexer, backend and frontend. A docker compose is provided so you don't have to stitch everything together yourself.

But livegrep is very well optimized, to the point that my current running instance against 200+ repos only consume 99 MB of memory for backend, and 21 MB for frontend, which is a rare sight in 2025.

The way livegrep works is that initially you'd run an indexer (it has direct integration with GitHub), which would clone the repos and generate the search index. Then you start a backend that picks up the said index, and frontend talks to backend via REST api.

On subsequent indexing, after an indexer job completes, you just have to restart the backend and frontend, respectively.

## New Kid on the Block

Recently there's `Sourcebot`, it uses the same innards as Sourcegraph but with different frontend implementation. It's slower than livegrep (very noticeable) and it doesn't work quite great when indexing 200+ repos. Which means....

## Some DIY is Needed

livegrep is perfect for everything except that it doesn't automatically clone your private repos by default. If private repos belong to an organization, there is no issues, but if they belong to a user, then livegrep can clone it, but you have to specify those repos manually. This is because the GitHub api livegrep uses to fetch your repos list does not return user's private repos (it's in the documentation).

To circumvent this, I wrote a small golang utility to fetch repos list (either belonging to orgs or users), then clone them. I also use Forgejo hosted on my home server as well, which means the default livegrep indexer doesn't support - going DIY means I can extend the repos cloning workflow to non-GitHub providers.

And the good thing is the code indexing workflow can be run separately in livegrep, so I just have to look up the code indexing config for livegrep, and add a step in my golang utility to generate it, then send it to livegrep indexer to generate the index and voila!

It's really satisfying when you can write code to solve that issue that's been bothering you.
