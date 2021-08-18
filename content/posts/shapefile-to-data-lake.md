---
title: Shapefile to data lake
date: 2021-04-23T18:25:13.000Z
draft: false
toc: false
images:
tags:
  - data engineering
  - gis
---


**Background**: we use spark to read/write to data lake. For dealing with spatial data & analysis, we use [sedona](http://sedona.apache.org). Shapefile is converted to TSV then read by spark for further processing & archival.

Recently I had to archive shapefiles in our data lake. It wasn't rosy for the following reasons:

## Invalid geometries

Sedona (and geopandas too) whines if it encounters `invalid geometry` during geometry casting. The invalid geometries could be from many reasons, one of them being unclean polygon clipping.

Solution: use `gdal` to filter out invalid geometries.

## Spatial projection

Geometric projections requires `projection`, otherwise you could be on the wrong side of the globe. This matters because by default, the worldwide-coverage projection is `EPSG:4326`, but the unit is in `degrees`, so sometimes for analysis the data is converted to a local projection which covers a smaller geographical region, but uses `meter` as the unit.

This means that if the source projection is in `A`, and you didn't cast it to `EPSG:4326`, spark would mistakenly think it's on `EPSG:4326` by default. Something like seeing the entirely of the UK in Africa.

Solution: verify the source projection and cast to `EPSG:4326` before writing to data lake.

## Extra new line character

Sometimes when editing shapefile data by hand using applications like ArcGIS or QGIS, you could copy a text which might contain "new line" character, and set it as a cell value. Spark doesn't play nice with "new line" characters in a middle of a record.

Solution: strip new line characters by hand.

Yes, I really did that 😶. Thankfully it was a very small shapefile that has the issue.

**Takeaways**: count yourself lucky if you never have to deal with spatial data.
