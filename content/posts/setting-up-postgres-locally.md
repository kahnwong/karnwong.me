---
title: Setting up Postgres locally, what could go wrong?
date: 2023-12-23T19:30:59+07:00
draft: false
ShowToc: false
images:
tags:
  - devops
  - database
  - networking
---

There are multiple reasons why someone wants to set up a postgres locally. Either for learning SQL or as an application's backend. Over the years I see people struggle with using postgres locally, so here are common use cases and possible issues, with solutions for each.

## For Learning SQL

SQL is very common for analysts to use for accessing data from a database, because the data size outgrows Excel. However, SQL is a query language, not a database engine. This essentially means if you want to get familiar with SQL, there are other simpler alternatives, such as SQLite or DuckDB (which can load data from local files directly without doing an explicit data import). Plus, you don't need authentication to use either of them!

This would save people a lot of headaches from fixing connection errors, which can be anything from wrong username/password, wrong host, host is not reachable, etc.

## As Application Backend

### Postgres setup

There are many ways to set up postgres locally. You can use an official installer to install postgres natively. This is very convenient, but you can only have one version of postgres installed at the same time.

However, given the version of postgres you use for your application is very important, I would recommend people to set up postgres via docker, since this allows you to choose any postgres version you want for your project, and you can reset the state of postgres easily, since you can mount the data to a host path.

### Networking issues

If you are using postgres via docker, and you are developing an application locally, don't forget to expose port `5432` when you start a postgres container, otherwise your application won't be able to discover postgres. You'll know it's working if you can discover this postgres via `localhost`.

But, if you are running your application as a docker container:

| Postgres container                                                           | Postgres hostname as seen from application |
| ---------------------------------------------------------------------------- | ------------------------------------------ |
| Expose port to `5432`                                                        | `host.docker.internal`                     |
| Don't expose port `5432` & run application and postgres via `docker compose` | `$postgresServiceName`                     |

Reason being from a docker container's perspective, it doesn't share `localhost` with your host system, which means `localhost` as seen on your machine, and the one seen in docker are completely different!

As for referencing a compose service name as database host, that's a docker networking feature, where each services in a single compose file can discover each other via a compose service name.

## Closing

Mostly this is networking issues, and not being familiar with docker networking model can trip most people up, hopefully this article is useful to you. Please let me know if you would like me to expand on some other use cases.
