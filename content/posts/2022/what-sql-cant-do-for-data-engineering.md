---
title: What SQL can't do for data engineering
date: 2022-05-15T22:38:45+07:00
draft: false
ShowToc: false
images:
tags:
  - data engineering

---

I often hear people ask "if you can do data engineering with SQL, then what's the point of learning spark or python?"

## Data ingestion

Let's circle back at bit.

I think we all can agree that: there's a point in time where there's no data in the data warehouse (which DE-who-use-SQL's use as base of operation). The source data could be anything from hard CSV/Excel or API endpoints. No data in data warehouse, DE can't use SQL to do stuff with the data.

So who put the data into the data warehouse? Data engineers of course! But not the kind that use SQL. These data engineers are known as "data _platform_ engineers," where their main focus is data ingestion, platform and scalability.

The confusion lies in the fact that, there seems to be different tools and skills required for platform or analytics type for data engineers, but some places still refer to both roles as data engineer.

And guess how data platform engineers write large amount of data to data warehouse. With help of spark and python/java of course!

## Data cleaning

What if you need to perform complex data cleaning processes for your data? You might have to create multiple temp columns, then coalesce them at the end to get the final result. You can do this in SQL, with the help of jinja template, to certain extent. But the same process expressed in python would be much more concise and readable. Especially if it in involves multiple steps spanning 700 lines of code. Not to mention you might need to perform debugging somewhere in between, and setting up SQL debugging is not something I wish on my worst enemy.

Some people say "but you can also create functions in SQL." Yes, you can, but it's clunky and very fragile, and not very readable.

## Optimization

When you're writing SQL for data transformation, the actual execution logic is being translated by the database's engine. As with all forms of translations, some information is lost, and can result in non-optimized instructions. For instance, if you need to perform longest matching against a list of string, one way to optimize this is to set a `break` condition after you found a match. In SQL, this most likely results in instructions to compare every string from the list, instead of stopping the matching process and move onto the next operation after a match is found.

---

These are some but not all instances where SQL falls short for data engineering tasks. I hope this article sheds some light on the importance of using python/spark for data engineering tasks.
