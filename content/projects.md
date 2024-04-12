---
title: "Projects"
layout: "projects"
url: "/projects"
summary: "projects"
ShowToc: false
---

## Ops

- [nix](https://www.karnwong.me/posts/2022/12/cross-platform-package-env-management-with-nix/) - A cross-platform setup script that works with both Linux and Mac.
- [self-hosted](https://github.com/kahnwong/self-hosted) - Self-hosting open-source alternatives for popular services. Managed via docker-compose, helm and terraform.
- [terraform-sops-ssm](https://github.com/kahnwong/terraform-sops-ssm) - Create SSM secrets from SOPS-encrypted secrets, with IAM roles & users creation for SSM access.
- [Vercel - Multi Branch Deployment](https://github.com/kahnwong/vercel-multi-branch-deployment) - Use GitHub Actions to deploy a frontend project from different branches (dev, uat, master), each with their own preview environment.
- [pgconn](https://github.com/kahnwong/pgconn) - pgcli wrapper to connect to PostgreSQL database specified in db.yaml. Proxy/tunnel connection is automatically created and killed when pgcli is exited.
- [totp](https://github.com/kahnwong/totp) - CLI TOTP token generator with autocomplete.
- [Calculator](https://calculator.karnwong.me/) - Various calculators: meeting, GKE Autopilot and cloud cost calculator. Source code [here](https://github.com/kahnwong/calculator).
- [Proxmox VM Selector](https://github.com/kahnwong/proxmox-vm-selector) - A simple TUI to select which Proxmox VM to start/stop.

## Data Engineering

- [Dataframe Frameworks Showdown](https://www.karnwong.me/posts/2023/04/duckdb-vs-polars-vs-spark/) - Benchmark performance between duckdb, polars and spark. In addition to runtime, RAM usage is also provided.
- [Spark on Kubernetes](https://www.karnwong.me/posts/2023/09/spark-on-kubernetes/)
- [Dagster Demo](https://github.com/kahnwong/dagster-demo)

## Data Science

- [Impute Pipelines](https://www.karnwong.me/posts/2020/05/impute-pipelines/) - Use machine learning to fill in missing data. Utilize hyperparameter tuning to find the optimum parameters.
- [Visualizing Map Region Prefix/Suffix](https://www.karnwong.me/posts/2020/09/visualizing-map-region-prefix-suffix/) - Utilize NLP to group region name prefix/suffix.
- [Word-Based Analysis With Song Lyrics](https://www.karnwong.me/posts/2020/04/word-based-analysis-with-song-lyrics/) - Visualize lyrics trend using NLP and use topic modeling to find common words per specified clusters.

## Tools

- [music-lyrics-tagger](https://github.com/kahnwong/music-lyrics-tagger) - Add lyrics to flac and m4a files.

## Automation

- [subsonic-github-readme](https://github.com/devbaygroup/subsonic-github-readme) - Now playing and random tracks widget via subsonic API. Golang port [here](https://github.com/kahnwong/subsonic-github-readme-golang).
- [todotxt-to-calendar](https://github.com/devbaygroup/todotxt-to-calendar) - Convert todo.txt entries to calendar all-day event.
- [water-cut-notify](https://github.com/kahnwong/water-cut-notify) - Send water cut alert as LINE notifications.
