+++
title = "Add Ghost content to Hugo"
date = "2021-03-31"
path = "/posts/2021/03/create-static-site-from-ghost-blog"

[taxonomies]
categories = [ "migration",]
tags = []

+++

Ghost CMS is very easy to use, but the deployment overhead (maintaining db, ghost version, updates and etc) might be too much for some. Luckily, there's a way to convert a Ghost site to static pages, which you can later host on Github pages or something similar.

## Setup

- static site engine: Hugo
- a Ghost instance

## Usage

1. Install [https://github.com/Fried-Chicken/ghost-static-site-generator](https://github.com/Fried-Chicken/ghost-static-site-generator)
2. `cd` to `static` directory in your Hugo folder
3. run

```bash
gssg --domain ${YOUR_GHOST_INSTANCE_URL} --dest posts --url ${YOUR_STATIC_SITE_DOMAIN_WITHOUT_TRAILING_SLASH} --subDir posts
```

4. Update your hugo config to link to the above folder:

```toml
[[menu.main]]
    identifier = "posts"
    name       = "Posts"
    url        = "/posts"
````

All done! 🎉🎉🎉
