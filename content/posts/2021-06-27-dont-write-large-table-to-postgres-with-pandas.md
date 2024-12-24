+++
title = "Don't write large table to Postgres with Pandas"
date = "2021-06-27"
path = "/posts/2021/06/dont-write-large-table-to-postgres-with-pandas"

[taxonomies]
categories = ["data-engineering",]
tags = ["database","postgres" ]

+++

We have a few tables where the data size is > 3GB (in parquet, so around 10 GB uncompressed). Loading it into postgres takes an hour. (Most of our tables are pretty small, hence the reason why we don't use columnar database).

I want to explore whether there's a faster way or not. The conclusion is writing to postgres with spark seems to be fastest, given we can't use `COPY` since our data contain free text, which means it would make CSV parsing impossible.

I also found out that the write performance from pandas to postgres is excruciatingly slow because:
It first decompresses the data in-memory. For a 30MB parquet (around 100MB uncompressed) it used more than 20GB of RAM (I killed the task before it finishes, since by this time the RAM usage is climbing up)

But even with reading plain JSON line in pandas with chunksize and use `to_sql` with `multi` option, it's still very slow.

In contrast, writing the said 30MB parquet file to postgres takes only 1 minute.

Big data is fun, said data scientists 🧪 (until they run out of RAM 😆)
