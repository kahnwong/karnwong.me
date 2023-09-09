---
title: Data engineer archtypes
date: 2022-08-26T10:06:36+07:00
draft: false
ShowToc: true
images:
tags:
  - data engineering
  - recommended
---

I have been working in the data industry since almost half a decade ago. Over time I have noticed so-called archetypes within various data engineering roles. Below are main skills and combinations I have seen over the years. This is by no means an exhaustive list, rather what I often see.

## SQL + SSIS

- Using SQL to manipulate data via SSIS, in which data engine is Microsoft SQL Server.
- Commonly found in enterprise organizations that use Microsoft stack.

## SQL + Hive

- Using SQL to manipulate data via Hive, a filesystem that support columnar data format, usually accessed via Zeppelin.
- Often found in enterprise organizations that work with big data before Spark was released.

## SQL + DBT

- Using SQL to manipulate data via DBT, an abstraction later for data pipelines scheduler that allows users to use SQL interface with various database engines. DBT is often mentioned in Modern Data Stack.
- Often found in newly established organizations in the last few years.

## Python + pandas

- Using python with pandas to manipulate data, usually with data that can fit into memory (ie less than 5GB)
- This is also common if you have data scientists manipulate data, since pandas is what they are familiar with. In addition, most people who write pandas are not known for writing well-optimized code, but it’s negligible for small data.

## Python + pyspark

- Using python with pyspark to manipulate data, can be either SQL or Spark SQL.
- Usually organizations that use pyspark also does machine learning as well.
- Often found in organizations that work with big data, and have established data lake platform.

## Scala + spark

- Using Scala to manipulate data via spark.
- Often found on enterprise organizations where they have been using spark before pyspark was released. Has more limited data ecosystem.

## Python + Task orchestrator (airflow, dagster, etc)

- Using task orchestrators to run pipelines on a regular basis, the application logic is written in python. Inside can be anything from pure python to pyspark. Or you can use bash and use any unix tools.
- People who fall under this category often have software engineering background.

## Platform engineering (setting up data infrastructure, etc)

- These are people that set up database, infrastructure, networking, and everything required to allow engineers/users to create data pipelines and consume data at downstream.
- Usually they are DevOps who transitioned from working with app infra to data infra.

---

**Updated 2022-09-02**

## GUI-based solutions

- Using GUI-based tools to create data pipelines, such as Talend, AWS Glue, Azure Data Factory, etc. May or may not use in conjunction with SQL / python / pyspark.
