+++
title = "Using Apache Iceberg to reduce data lake operations overhead"
date = "2023-11-15"
path = "/posts/2023/11/using-apache-iceberg-to-reduce-data-lake-operations-overhead"

[taxonomies]
categories = [ "data-engineering",]
tags = [ "data-lake", "iceberg", "spark",]

+++

Every business generates data, some very little, some do generate ginormous amount of data. If you are familiar with the basic web application architecture, there are data, application and web tier. But it doesn't end there, because the data generated has to be analyzed for reports. A lot of organizations have analysts working on production database directly. This works fine and well, until the data they are working with is very large to the point that a single query can take half a day to process!

So what do people do? This is when data engineers set up pipelines to move data from $sourceDataLocation to $destinationDataLocation. Sounds simple enough, but the data doesn't only move once! Because most often than not, data from $sourceDataLocation is not in a usable format by analysts. Some transformations have to be made, columns have to be added or joined with other related data, etc. This process is known as Extract, Transform, Load (ETL, but these days you should aim for ELTLTLTLTLTLTLTLTL.)

The difference between doing a Load before a Transform is that, if your transformation tasks failed, you don't have to re-ingest the source data! At small data volume this would add more overhead, but if we are talking about 100GB then this process adds more resilience to the pipeline architecture.

Since this is big data territory, the general consensus is to use the data lake model, where data are stored in a blob storage, since this means cheaper cost (compared to storing data in a database, either relational or columnar), and more throughput for writing and accessing the data. Spark is used for data processing, since it is a distributed dataframe framework.

As with the nature of businesses, some changes would be made to the web application along the way, either due to more features added, or a direction shift. This means over time, there would be new data columns added, some renamed to better fit the purpose of a business.

However, this does pose a few issues, namely accessing the data is not as convenient (because it involves a user specifying actual data path in blob storage), and if you want to rename or add columns, you would have to manually go through existing data and update it in the data lake manually. For orgs that have petabyte-scale of data, a single schema change can mean a few weeks of backfilling pipelines.

Non-intuitive data access pattern and the need for schema evolution are very painful experiences for data engineers, and they eat up a lot of engineering time to make sure data lake operations are running smoothly. I discovered Apache Iceberg last year, was recently tinkered with it, and it solves a lot of major data lake operation headaches in following ways (there are many more, so check out their [official docs!](https://iceberg.apache.org/docs/latest/))

| Area | Spark | Iceberg |
| ---- | -----| ----------------- |
|Parquet file size | Have to manually optimize so that each resulting parquet would be around 128MB for optimum performance. | Automatically taken care of. |
| Schema evolution | Have to re-process existing data so all files would have the same schema, otherwise it would break a .read() operation. | Can execute SQL DDL without requiring full table rewrite. |
| Versioning | Have to utilize blob storage versioning, but to actually revert back to a previous state still requires a lot of configuration. | Can freely time-travel to the desire revision, since all writes to Iceberg are stored. This mean a table overwrite doesn't actually delete existing data, it is an append underneath, just that the presentation layer only show the latest data state. |
| Partitioned table | Have to create a partition-key column. After writing a table, partition keys can't be updated, unless you perform a full table rewrite. | Can synthesize a partition column on-the-fly (ex. use timestamp column as partition key, but set it as date), with support for partition key modifications after initial table write. |
| Add new table partition | Have to configure spark so it doesn't delete existing partitions if they are not present in your current dataframe. | Iceberg supports partition key in the DDL, and on table write you can specify whether it is an overwrite or overwrite on partition.|

Also technically, Iceberg doesn't require a stateful component if you utilize hadoop catalog (think of a manifest of all table revisions), because the manifests are stored together with the data as normal json files. However, this means concurrent writes are not possible. To prevent this, you can select available [Iceberg catalogs](https://iceberg.apache.org/concepts/catalog/). For a lightweight and vendor-agnostic catalog, you can use REST catalog with Postgres, and these would be only required stateful backend for Apache Iceberg setup.

Did I mention that by default, Iceberg uses ztsd compression, which means smaller file size than the default snappy compression on parquet?
