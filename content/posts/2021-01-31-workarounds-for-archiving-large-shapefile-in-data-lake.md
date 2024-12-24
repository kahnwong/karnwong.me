+++
title = "Workarounds for archiving large shapefile in data lake"
date = "2021-01-31"
path = "/posts/2021/01/workarounds-for-archiving-large-shapefile-in-data-lake"

[taxonomies]
categories = [ "data-engineering",]
tags = [ "gis",]

+++

If you work with spatial data, chances are you are familiar with `shapefile`, a file format for viewing / editing spatial data.

Essentially, shapefile is just a tabular data like `csv`, but it does thingamajig with `geometry` data type, where any gis tools like `qgis` or `arcgis` can understand right away. If you have a csv file with geometry column in `wkt` format (something like `POLYGON ((30 10, 40 40, 20 40, 10 20, 30 10))`), you'll have to specify which column is to be used for geometry.

If you want to store shapefile in data lake, it's best to store it as parquet or any format you normally use, since it's faster to read and filter. For comparison, parsing a 5GB+ shapefile and filter takes _longer_ than reading a gzipped json, filter, and export to shapefile.

Normally I would use `geopandas` to read spatial data and convert it to pandas dataframe, then send it to spark. But since the shapefile is very large, it takes forever to read in geopandas. This tells me that there is a parsing bottleneck going on. And geopandas can't read shapefile with multiple geometry types (this shouldn't happen, but sometimes during editing, clipping this here and there can cause invalid geometry).

Qgis has a tool to fix invalid geometries, so I tried exporting shapefile to csv, but qgis went OOM. But both qgis and geopandas use `gdal` for backend, and it has a CLI interface, so I look up how to export shapefile to tsv (`tab` as a seperator makes it faster to parse since it rarely occurs).

Now things work perfectly. As a bonus, gdal also skip invalid geometries by default (unlike in geopandas where it will throw an error and there's no way to ignore it and tell the parser to keep going).

At this point I have a nice tsv file, and reading & archiving via spark is now a breeze. Yay!

## Takeaway

- If it takes too long to read, maybe it's a parsing bottleneck. Find a way to convert it to another format so it's easier to parse.
- Sometimes your initial tools of choice might have some quirks. In most cases there will be similar tools out there that can workaround the issues. (In this case, use gdal to convert to csv in lieu of geopandas because gpd can't work with invalid geometries & takes longer to read compared to feeding spark a straight csv/tsv).
