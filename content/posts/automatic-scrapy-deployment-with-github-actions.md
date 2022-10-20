---
title: "Automatic scrapy deployment with GitHub actions"
date: 2021-06-02T21:48:12+07:00
draft: false
ShowToc: false
images:
tags:
  - devops
---

Repo [here](https://github.com/kahnwong/scrapy-deploy-gh-actions)

Scrapy is a nice framework for web scraping. But like all local development processes, some settings / configs are disabled.

This wouldn't pose an issue, but to deploy a scrapy project to zyte (a hosted scrapy platform) you need to run `shub deploy`, and if you run it and forget to reset the config back to prod settings, a Titan may devour your home.

You can set auto deployment from github via the UI in zyte, but it only works with github only. Plus if you want to run some extra tests during CI/CD you're out of luck. So here's how to set up CI/CD to deploy automatically:

**Note**: I would assume that you have your scrapy project set up already.

## Create scrapinghub.yml + add repo secrets

```yaml
project: ${PROJECT_ID}

requirements:
  file: requirements.txt

stack: scrapy:${YOUR_SCRAPY_VERSION_IN_PIPFILE}
apikey: null
```

Notice that `apikey` is left blank. This is because it's considered a good practice to not check in sensitive information & credentials in version control. Instead `apikey` will be added to github secrets, so it can be called as environment variable.

## Create github workflow file

```yaml
name: Deploy

on:
  push:
    branches: [master, main]
jobs:
  deploy:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - name: Set up Python 3.9
        uses: actions/setup-python@v2
        with:
          python-version: 3.9
      - name: Install dependencies
        run: |
          python -m pip install --upgrade pip
          pip install pyyaml shub
      - name: Deploy to zyte
        if: github.ref == 'refs/heads/master'
        run: python3 utils/edit_deploy_config.py && shub deploy
        env:
          APIKEY: ${{ secrets.APIKEY }}
```

Translation:

- On push to this repo (this doesn't work for PRs)
- Download this repo
- Setup python3.9
- Install some pip modules
- Run a script to overwrite scrapinghub.yml's apikey value, in which the value is obtained from github secrets
- Execute deploy command
