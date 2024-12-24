+++
title = "Password auth with apache2 reverse-proxy"
date = "2021-02-22"
path = "/posts/2021/02/setting-up-password-auth-with-apache2-reverse-proxy"

[taxonomies]
categories = ["web-hosting",]
tags = [ ]

+++

EDIT: see [here](/posts/2021/03/hello-caddy) for Caddy, also easier to set up too.

Sometimes you found an interesting project to self-hosted, but it doesn't have password authentication built-in. Luckily, we need to reverse-proxy them anyway and apache2/ nginx / httpd happen to provide password auth with reverse-proxy by default.

To set up password auth with apache2 via reverse-proxy:

1. `echo "${PASSWORD}" | htpasswd -c -i /etc/apache2/.htpasswd ${USER}` on your host machine which has apache2 installed.
2. create a vhost config:

```xml
<VirtualHost *:80>
    ProxyPreserveHost On

    ProxyPass / http://localhost:${EXPOSED_CONTAINER_PORT}/
    ProxyPassReverse / http://localhost:${EXPOSED_CONTAINER_PORT}/

    ServerName ${YOUR_DOMAIN}

    <Proxy *>
        Order deny,allow
        Allow from all
        Authtype Basic
        Authname "Password Required"
        AuthUserFile /etc/apache2/.htpasswd
        Require valid-user
    </Proxy>
</virtualhost>
```

That's it!
