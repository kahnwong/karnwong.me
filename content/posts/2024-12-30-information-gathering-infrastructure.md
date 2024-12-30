+++
date = "2024-12-30"
path = "/posts/2024/12/information-gathering-infrastructure"
title = "Information gathering infrastructure"

[taxonomies]
categories = ["homelab"]
tags = ["automation"]

[extra]
mermaid = true
+++

{% mermaid() %}
flowchart LR

Feed --> Miniflux

%% feeds
subgraph Homelab
Miniflux --> Wallabag
end

Wallabag -->KOReader

%% newsletters
subgraph SaaS
Newsletters --> Email
end

subgraph Automation
Email --> EmailToEpubScript
end

EmailToEpubScript --> Syncthing

%% device
subgraph Device
KOReader --> Ereader
Syncthing --> Ereader
end
{% end %}

Once upon a time, I was asked by a recruiter regarding how I keep up with news and industry trends. I described something along this diagram, but apparently it went over his head.

I'm an avid reader, but reading on a computer or tablet screen is not fun for prolonged period of time. I've been using an ereader since 2012, and I love how easy it is on my eyes.

Back then things were simple, almost every site provided RSS. Then Google Reader got shut down, and people moved onto newsletters model, although some sites still provide RSS. Since I prefer to read stuff on my ereader, the flow is like this:

## For feeds

1. Subscribe them via Miniflux, an RSS feed aggregator. It's hosted on my homelab.
2. For articles I want to read later, I send them to Wallabag via Miniflux integration. Walalbag is a Pocket-like app, self-hostable.
3. On my ereader, I have KOReader installed. It has a Wallabag plugin to download and sync articles.

## For newsletters

1. Use an email address for newsletter subscription to subscribe newsletters.
2. Use a Go program to fetch emails and convert them to epub
3. Put epub files in a Syncthing folder. Syncthing is a file-sync application, this should be on both your workstation and ereader
4. On ereader, sync a folder that you use via Syncthing to obtain epub files.
