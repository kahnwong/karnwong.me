+++
title = "Mongodb export woes"
date = "2021-01-27"
path = "/posts/2021/01/mongodb-export-woes"

[taxonomies]
categories = ["data-engineering",]
tags = [ "database" ]

+++

There's a task where I need to export 4M+ records out of mongodb, total uncompressed size is ~~17GB+~~ 26GB

## export methods

### mongoexport

The recommended way to export is using `mongoexport` utility, but you have to specify the output attributes, which doesn't work for me because the schema from older set of records are less than the newer set

### DIY python script

#### the vanilla way

But you can interact with mongodb from python, and if you read from it it'll return a dict, which is perfect for this because you don't have to specify the required attributes beforehand. So what I do is:

```python
cursor = collection.find({})
total_records = collection.estimated_document_count()

with open(filename, "w") as f:
    for i in tqdm(cursor, total=total_records):
        f.write(json.dumps(i, default=myconverter, ensure_ascii=False))
        f.write("\n")
```

The cons for this solution is it needs a lot of hdd space since it's uncompressed. But it **works best if you need to export a collection with mismatched schema**.

#### the incremental export way

You can also incrementally export your collection from mongodb using `.skip($START_INDEX).limit($INCREMENT_SIZE)` , but it **performs worse than the vanilla way**, since what mongodb does is just iterating through everything all over again to get to your specified `start:end` index.

## Performance comparison

On my local machine (<10 MB/s transfer speed) I could export a collection with around 4.5M records within 1 hour, but on a VPS with incremental export it takes 9 hours and counting.

## Takeaway

Please do not store a large dataset in mongodb where you need to dump everything out, especially if you use it as a raw data source. It's fine if you store prepped output for API to be queried via `_id` (primary key).
