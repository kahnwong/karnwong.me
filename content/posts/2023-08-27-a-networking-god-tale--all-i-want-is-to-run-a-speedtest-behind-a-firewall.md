+++
title = "A networking God tale: all I want is to run a speedtest behind a firewall"
date = "2023-08-27"
path = "/posts/2023/08/a-networking-god-tale--all-i-want-is-to-run-a-speedtest-behind-a-firewall"

[taxonomies]
categories = ["infrastructure"]
tags = [ "networking",]

+++

Imagine going to your client's site to deploy a software. During the deployment process, you notice that the speed is atrociously slow. You have a suspicion that your client's network bandwidth is the issue.

To test this theory, you go to a speedtest website and run a test. Turns out you can't because it's blocked at the firewall level. Then you try another speedtest website, oops still got blocked. Then you try a few more, still no dice.

At this point you're pretty sure that your client downloaded a hostfile blocklist and apply it to their firewall. But you need to know that the issue is at your client's network bandwidth, so you can tell them with confidence that the issue is them, not you.

Stepping back a bit, a speedtest is just network performance on download and upload. So technically if you have a server that allows you to download and upload files, you can also measure the transfer speed.

With this insight, you spin up a cheap VPS, then deploy a file sharing server. On your client's machine, you download a large binary from somewhere, maybe a linux ISO installer image. Then you upload it to your VPS. You write down the upload speed. Then you download the same file from the VPS, this should give you the download speed.

Now you have what you need to tell your client that the issue is on their end!

Maybe not really a happy ending, but at least you know your code works fine (for now).
