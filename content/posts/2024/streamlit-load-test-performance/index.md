---
title: Streamlit load test performance
date: 2024-09-07T16:12:17+07:00
draft: false
ShowToc: false
images:
tags:
  - python
---

Streamlit is well-loved by many people, especially among data folks due to the fact that it does not require prior web programming knowledge to get started.

Popular use cases for streamlit can be anything from a quick machine learning application poc or internal dashboards. But what if you want to create a production deployment? Would streamlit still be a viable option?

## Experiment setup

For a production deployment, let's assume there are 500 concurrent users. For a single-page streamlit with a chat box taking 3 seconds to return a single-paragraph message:

![streamlit-ui.webp](images/streamlit-ui.webp)

With following configuration for the load test:

```yaml
config:
  phases:
    - duration: 100
      arrivalRate: 50
      maxVusers: 500
```

## Benchmark result

```bash
--------------------------------
Summary report @ 12:50:35(+0700)
--------------------------------

browser.http_requests: ................................................... 5926
browser.page.TTFB.http://localhost:8501/:
  min: ................................................................... 3
  max: ................................................................... 19934.3
  mean: .................................................................. 6975.6
  median: ................................................................ 7260.8
  p95: ................................................................... 16486.1
  p99: ................................................................... 18963.6
browser.page.codes.200: .................................................. 5950
errors.page.goto: Timeout 30000ms exceeded.: ............................. 1486
errors.page.goto: net::ERR_CONNECTION_RESET at http://localhost:8501/: ... 9
vusers.completed: ........................................................ 421
vusers.created: .......................................................... 1916
vusers.created_by_name.0: ................................................ 1916
vusers.failed: ........................................................... 1495
vusers.session_length:
  min: ................................................................... 1874.4
  max: ................................................................... 30379.4
  mean: .................................................................. 23027.7
  median: ................................................................ 25598.5
  p95: ................................................................... 30040.3
  p99: ................................................................... 30040.3
vusers.skipped: .......................................................... 3084
```

Notice the `median` value of `25598.5`, this is roughly around `25 seconds` to render the above streamlit page.

## Takeaway

Moral of the story: streamlit is great for a quick poc or low traffic apps. Anything more than this and you should consider using a dedicated web framework to create the frontend.

Also, yes, CPU and memory usage also shot up.
