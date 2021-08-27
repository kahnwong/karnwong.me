---
title: "Data engineering toolset (that I use) glossary"
date: 2021-06-04T23:57:58+07:00
draft: false
toc: false
images:
tags:
  - data engineering
---

# Big data
- Spark: Map-reduce framework for dealing with big data, especially for data that doesn't fit into memory. Utilizes parallelization.

# Cloud
- AWS: Cloud platform for many tools used in software engineering.
- AWS Fargate: A task launch mode for `ECS task`, where it automatically shuts down once a container exits. With EC2 launch mode, you'll have to turn off the machine yourself.
- AWS Lambda: Serverless function, can be used with docker image too. Can also hook this with API gateway to make it act as API endpoint.
- AWS RDS: Managed databases from AWS.
- ECS Task: Launch a task in ECS cluster. For long-running services, launch via EC2. For small periodical tasks, trigger via Cloudwatch. For the latter, think of cron-like schedule for a task. Essentially at specified time, it runs a predefined docker image (you should configure your `entrypoint.sh` accordingly).

# Data
- Parquet: Columnar data blob format, very efficient due to column-based compression with schema definition baked in.

# Data engineering
- Dagster: Task orchestration framework with built-in pipelines validatioin.
- ETL: Stands for extract-transform-load. Essentially it means "moving data from A to B, with optional data wrangling in the middle."

# Data science
- NLP: Using machine (computer) to work on human languages. For instance, analyze whether a message is positive or negative.

# Data wrangling
- Pandas: Dataframe wrangler, think of programmable Excel.
# Database
- Postgres: RMDBS with good performance.
# DataOps
- Great expectations: A framework for data validation.

# DevOps
- Docker: Virtualization via containers.
- Git: Version control.
- Kubernetes: Container orchestration system.
- Terraform: Infrastructure as code tool, essentially you use it to store a blueprint for your infra setup. If you were to move to another account, you can re-conjure existing infra with one command. This makes editing infra config easier too, since it automatically cleans up / update config automatically.

# GIS
- PostGIS: GIS extension for Postgres.

# MLOps
- MLflow: A framework to track model parameters and output. Can also store model artifact as well.

# Notebook
- Jupyter: Python notebook, used for exploring solutions before converting it to .py.
