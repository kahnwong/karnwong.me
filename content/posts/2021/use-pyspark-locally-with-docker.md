---
title: Use pyspark locally with docker
date: 2021-12-21T19:26:32+07:00
draft: false
ShowToc: false
images:
tags:
  - devops
  - data science
  - data engineering
  - docker

---

For data that doesn't fit into memory, spark is often a recommended solution, since it can utilize map-reduce to work with data in a distributed manner. However, setting up local spark development from scratch involves multiple steps, and definitely not for a faint of heart. Thankfully using docker means you can skip a lot of steps 😃

## Instructions

1. Install [Docker Desktop](https://www.docker.com/get-started)
2. Create `docker-compose.yml` in a directory somewhere

```yml
version: "3.3"

services:
  pyspark:
    container_name: pyspark
    image: jupyter/pyspark-notebook:latest
    ports:
      - "8888:8888"
    volumes:
      - ./:/home/jovyan/work
```

3. Run `docker-compose up` from the same folder where the above file is located.
4. You should see something like this. It's the same from running `jupyter notebook` locally. Click the link at the end to access jupyter notebook.

```log
Creating pyspark ... done
Attaching to pyspark
pyspark    | WARNING: Jupyter Notebook deprecation notice https://github.com/jupyter/docker-stacks#jupyter-notebook-deprecation-notice.
pyspark    | Entered start.sh with args: jupyter notebook
pyspark    | /usr/local/bin/start.sh: running hooks in /usr/local/bin/before-notebook.d as uid / gid: 1000 / 100
pyspark    | /usr/local/bin/start.sh: running script /usr/local/bin/before-notebook.d/spark-config.sh
pyspark    | /usr/local/bin/start.sh: done running hooks in /usr/local/bin/before-notebook.d
pyspark    | Executing the command: jupyter notebook
pyspark    | [I 12:36:04.395 NotebookApp] Writing notebook server cookie secret to /home/jovyan/.local/share/jupyter/runtime/notebook_cookie_secret
pyspark    | [W 2021-12-21 12:36:05.487 LabApp] 'ip' has moved from NotebookApp to ServerApp. This config will be passed to ServerApp. Be sure to update your config before our next release.
pyspark    | [W 2021-12-21 12:36:05.488 LabApp] 'port' has moved from NotebookApp to ServerApp. This config will be passed to ServerApp. Be sure to update your config before our next release.
pyspark    | [W 2021-12-21 12:36:05.488 LabApp] 'port' has moved from NotebookApp to ServerApp. This config will be passed to ServerApp. Be sure to update your config before our next release.
pyspark    | [W 2021-12-21 12:36:05.488 LabApp] 'port' has moved from NotebookApp to ServerApp. This config will be passed to ServerApp. Be sure to update your config before our next release.
pyspark    | [I 2021-12-21 12:36:05.497 LabApp] JupyterLab extension loaded from /opt/conda/lib/python3.9/site-packages/jupyterlab
pyspark    | [I 2021-12-21 12:36:05.498 LabApp] JupyterLab application directory is /opt/conda/share/jupyter/lab
pyspark    | [I 12:36:05.504 NotebookApp] Serving notebooks from local directory: /home/jovyan
pyspark    | [I 12:36:05.504 NotebookApp] Jupyter Notebook 6.4.6 is running at:
pyspark    | [I 12:36:05.504 NotebookApp] http://bd20652c22d3:8888/?token=408f2020435dfb489c8d2710736a83f7a3c7cd969b3a1629
pyspark    | [I 12:36:05.504 NotebookApp]  or http://127.0.0.1:8888/?token=408f2020435dfb489c8d2710736a83f7a3c7cd969b3a1629
pyspark    | [I 12:36:05.504 NotebookApp] Use Control-C to stop this server and shut down all kernels (twice to skip confirmation).
pyspark    | [C 12:36:05.509 NotebookApp]
pyspark    |
pyspark    |     To access the notebook, open this file in a browser:
pyspark    |         file:///home/jovyan/.local/share/jupyter/runtime/nbserver-7-open.html
pyspark    |     Or copy and paste one of these URLs:
pyspark    |         http://bd20652c22d3:8888/?token=408f2020435dfb489c8d2710736a83f7a3c7cd969b3a1629
pyspark    |      or http://127.0.0.1:8888/?token=408f2020435dfb489c8d2710736a83f7a3c7cd969b3a1629
```

---

This snippet

```yml
volumes:
  - ./:/home/jovyan/work
```

means that anything you put in [the folder where `docker-compose.yml` is] can be accessed by [jupyter notebook running inside docker], and what you do from inside jupyter notebook can be seen on the host system too.

See? Easy as a 🥧.
