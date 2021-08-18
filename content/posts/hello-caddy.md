---
title: Hello Caddy
date: 2021-03-07T08:32:19.000Z
draft: false
toc: false
images:
tags:
  - infra
---

Since starting self-hosting back in 2017, I've always used apache2 since it's the first webserver I came across. Over time adding more services and managing separate vhost config is a bit tiresome.

Enters Caddy. It's very simple to set up and configure. Some services where I have trouble setting up in apache2 do not need extra config at all, even TLS is set up by default. Starting from Caddy2 it works with CNAME by default without extra setups.

You can set it up using a Caddy docker container, but some containers I use also expose port 443, so I have to install Caddy natively instead. 

For multiple sites config setup:

```caddyfile
# /etc/caddy/Caddyfile

SUBDOMAIN1.DOMAIN.com {
    reverse_proxy 127.0.0.1:${PORT}
}
SUBDOMAIN2.DOMAIN.com {
    reverse_proxy 127.0.0.1:${PORT}
}
```

For basic authentication, it's very, very simple (to the point I regret time researching it in apache2):

```caddyfile
# generate password hash
caddy hash-password --algorithm bcrypt

# add basicauth to Caddyfile
SUBDOMAIN1.DOMAIN.com {
    basicauth * {
        ${USERNAME} ${CADDY_PASSWORD_HASH}
    }
    reverse_proxy 127.0.0.1:${PORT}
}
```

And run `systemctl reload caddy`. You're all set!