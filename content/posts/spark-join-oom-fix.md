---
title: Spark join OOM fix
date: 2021-04-11T16:20:23.000Z
draft: false
ShowToc: false
images:
tags:
  - data
---
I have a big pipelines where one step performs crossjoin on `130K x 7K`. It fails quite often, and I have to pray to the Rice God for it to pass. Today I found the solution: `repartition` before crossjoin.

The root cause is that the dataframe with 130K records has 6 partitions, so when I perform crossjoin (one-to-many) it's working against those 6 partitions. Total output in parquet is around 350MB, which means my computer (8 cores, 10GB RAM provisioned for spark) needs to be able to hold all uncompressed data in memory. It couldn't hence the frequent OOM.

So by increasing the partition size from 6 to 24, the current working dataframe size is smaller, which means things could pass along faster while not filling up my machine's RAM.
