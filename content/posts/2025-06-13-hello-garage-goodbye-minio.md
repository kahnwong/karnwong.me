+++
title = 'Hello Garage, goodbye MinIO'
date = '2025-06-13'
path = '/posts/2025/06/hello-garage-goodbye-minio'

[taxonomies]
categories = ['infrastructure']
tags = []
+++

Given the trend where open-source software license changes to be less permissive, or move features behind a paywall,
existing users might not be quite as happy about this predicament. Case in point: [MinIO Removes Web UI Features from Community Version, Pushes Users to Paid Plans](https://news.ycombinator.com/item?id=44136108).

But this post is not about that, but rather more on why I'm migrating to garage.

Recently I've been using minio as a dvc backend (think of git LFS but using blob storage as backend) and it accumulated
a lot of small files. Understandably, minio performance tanked because it doesn't handle small files well. Also minio
deployment seems to use a little too much resources - it consumes 218 MB of memory. This much memory usage would be expected of a web application, but not for something like minio. Also see [this minio issue](https://github.com/minio/minio/issues/9966).

To get started with garage, follow this [quickstart](https://garagehq.deuxfleurs.fr/documentation/quick-start/). Adjust the
toml config as necessary.

Note that aws s3 api performs a checksum, and recent aws sdk broke garage's default settings.
See [this issue](https://github.com/boto/boto3/issues/4392#issuecomment-2868118431).

In essence, to get the most recent of garage (v1.1.0) working with aws cli and sdk, following configurations should be
made:

#### aws config

- set region to the value you specified in `garage.toml`

#### environment variables

For preventing invalid checksum errors.

```bash
AWS_REQUEST_CHECKSUM_CALCULATION=when_required
AWS_RESPONSE_CHECKSUM_VALIDATION=when_required
```

## Resources consumption

I didn't expect this much difference, but it's very welcome.

| Service | CPU | Memory |
|---------|-----|--------|
| MinIO   | 3   | 218    |
| Garage  | 1   | 5      |

Because garage is written in rust, score for Ferris the Crab!
