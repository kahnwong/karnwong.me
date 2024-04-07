---
title: Slim down python docker image size with poetry and pip
date: 2024-04-07T13:27:48+07:00
draft: false
ShowToc: false
images:
tags:
  - python
  - docker
  - devops
---

Python package management is not straightforward, seeing default package manager (pip) does not behave like node's npm, in a sense that it doesn't track dependencies versions.

This is why you should use `poetry` to manage python packages, since it creates a lock file, so you can be sure that on every re-install, the versions would be the same.

However, this poses a challenge when you want to create a docker image with poetry, because you need to do an extra `pip install poetry` (unless you bake this into your base python image). Additionally, turns out using poetry to install packages results in larger docker image size.

## Dockerfiles

Below are dockerfiles I use to compare between using `poetry` and `pip`:

### Dockerfile.poetry

```Dockerfile
FROM python:3.12-slim

WORKDIR /app

# hadolint ignore=DL3013
RUN pip install --no-cache-dir poetry

COPY pyproject.toml .
COPY poetry.lock .

RUN poetry install --only main --no-root --no-directory
```

### Dockerfile.pip

```Dockerfile
FROM python:3.12-slim

WORKDIR /app

COPY requirements.txt .
RUN pip install -r requirements.txt --no-cache-dir
```

## Installed packages

```toml
requests = "^2.31.0"
polars = "^0.20.18"
fastapi = "^0.110.1"
pydantic = "^2.6.4"
python-dotenv = "^1.0.1"
langchain = "^0.1.14"
psycopg2-binary = "^2.9.9"
```

## Result

And the resulting image size is:

```log
benchmark_poetry      latest            23d3105ad0dd   11 seconds ago   520MB
benchmark_pip         latest            b7932a02a8d1   12 hours ago     388MB
```

As you can see, using poetry makes the image `132 MB` larger. Let's say you deploy 12 times per month, that's extra `1584 MB`.

While I agree that these days storage is cheap, reducing images size here and there won't hurt 😎.
