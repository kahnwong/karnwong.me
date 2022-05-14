---
title: Reduce docker image size with alpine
date: 2021-12-19T20:50:07+07:00
draft: false
ShowToc: false
images:
tags:
  - devops
---

Creating scripts are easy. But creating a small docker image is not 😅.

Not all Linux flavors are created equal, some are bigger than others, etc. But this difference is very crucial when it comes to reducing docker image size.

## A simple bash script docker image
Given a Dockerfile (change `apk` to `apt` for `ubuntu`):
```dockerfile
FROM alpine:3

WORKDIR /app

RUN apk update && apk add jq curl

COPY water-cut-notify.sh ./

ENTRYPOINT ["sh", "/app/water-cut-notify.sh"]
```

| Base image | Docker image size |
| ---------- | ----------------- |
| alpine     | 11.1MB            |
| ubuntu     | 122MB             |

Ubuntu imag size is `1099%` larger!!!!!!

## What about a light python image?
```dockerfile
FROM python:3.9-alpine

WORKDIR /app

RUN pip install requests
```

| Base image | Docker image size |
| ---------- | ----------------- |
| alpine     | 53.6MB            |
| ubuntu     | 920MB             |

Ubuntu imag size is `1716%` larger!!!!!!



## Should you use alpine image for everything?
From above two experiments, we get:

![](/images/2021-12-19-21-21-35.png)

It's obvious that alpine has a very significant lighter footprint than ubuntu. But don't use alpine image for everything. For bash and go, using alpine results in lighter footprint. But for python apps, it's better to go with debian-based images. This [article](https://pythonspeed.com/articles/alpine-docker-python/) explains in details why it's so.
