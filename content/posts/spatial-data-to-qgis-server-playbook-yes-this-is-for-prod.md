---
title: Spatial data to QGIS server playbook (yes, this is for prod)
date: 2023-08-10T21:32:59+07:00
draft: false
ShowToc: false
images:
tags:
  - gis
---

Some of you might be familiar with geoserver for serving spatial data as consumable WMS/WFS layers. The issue is that as far as I know, you have to manually manage assets upload and specifying styles manually. Also the tool is a bit dated.

One modern alternative is QGIS server, you can find pre-made docker image online, and it also syncs with the Desktop version. The good thing about QGIS server is that you can create a QGIS project via the desktop application, then upload it wholesale to Postgres instance as QGIS server backend.

Assuming we're starting with raw spatial data as hardfiles, and you want to use Postgres as data store backend, and that you don't want to expose the database to public:

## Playbook

1. Import data to QGIS desktop.
2. `Fix geometries` > save to Postgres.
3. `Collect geometries` > save to Postgres. (This is for some cases where the raw data contain both `Polygon` and `MultiPolygon` in the same file, Postgres backend doesn't allow mixed geometry type).
4. `DB Manager` > `Import Layer/File`, and set desired `CRS` + spatial index here.
5. Import uploaded layers into current QGIS project.
6. Import geoserver styles (in SLD format) in QGIS via python console, then attach to layers.

```python
layer = iface.activeLayer()

layer.loadSldStyle("style.sld")
```

7. `Properties` > set `rendering` to `3`.
8. Enable `WMS` and `WFS` under `project properties`.
9. Save current project as `*.qgz`.
10. Unzip `*.qgz`, manually replace `database hostname` with `private ip address`.
11. Re-zip the assets into `*.qgz`. Important: have to use command-line on MacOS.
12. Turn on airplane mode, then open the modified QGIS project. (This is so that it won't keep reaching Postgres via private ip.)
13. Disable airplane mode, then write the project to Postgres. (Need to whitelist your home ip so that you can reach prod db behind vpc.)
14. Voila 🎉

## Notes

- To be sure, restart the server every time you update the project.
- Need to use private ip for database hostname.
- On M1, need to use QGIS installed via mamba. (Otherwise you'll have missing requirements for `Fix geometries`).
