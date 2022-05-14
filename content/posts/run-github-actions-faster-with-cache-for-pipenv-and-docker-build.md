---
title: Run GitHub Actions faster with cache for pipenv and docker build
date: 2021-11-09T23:54:23+07:00
draft: false
ShowToc: false
images:
tags:
  - devops
  - github
---

_Update 2021-11-29_

Recently we create more PRs, notice that there are a lot of redundant steps (env setup before triggering checks, etc). Found out you can cache steps in GitHub Actions, so I did some research. Got it working and turns out I reduce at least 60% actions time for a large docker image build (since only the later `RUN` directives are changed more frequently). For pipenv it shaved off 1 minute 18 seconds. Pretty neat!

### pipenv cache
```yaml
# https://github.com/actions/setup-python#caching-packages-dependencies
- uses: actions/setup-python@v2
  with:
    python-version: '3.9'
    cache: 'pipenv'
```

### docker build cache
```yaml
- name: Set up Docker Buildx
  uses: docker/setup-buildx-action@v1

- name: Add IMAGE_TAG to env
  run: echo "IMAGE_TAG=`echo ${GITHUB_SHA::7}`" >> $GITHUB_ENV

- name: Build, tag, and push image to Amazon ECR
  uses: docker/build-push-action@v2
  env:
    IMAGE_NAME: service/app
  with:
    context: .
    push: true
    tags: ${{ env.IMAGE_NAME }}:latest,${{ env.IMAGE_NAME }}:${{ env.IMAGE_TAG }}
    cache-from: type=gha
    cache-to: type=gha,mode=max
```
