+++
date = "2025-03-02"
path = "/posts/2025/03/migrate-from-docker-desktop-to-podman-on-darwin"
title = "Migrate from Docker Desktop to Podman on Darwin"

[taxonomies]
categories = ["devops"]
tags = ["darwin", "docker"]
+++

Lately Docker Desktop on Mac behaves weirdly, namely its erratic memory consumption. Not willing to babysit the memory usage, I looked into alternatives, and found a few contenders: Colima, Podman and Orbstack.

Being an open-source aficionado, naturally I gravitate to Colima and Podman. Although my spidey sense tells me Podman is more polished. Gave it a go and it works beautifully, so here's my setup to make it backward compatible with tools utilizing docker command and socket.

## Installing Podman

```bash
brew install --cask podman-desktop
brew install podman
brew install podman-compose
```

## Shell Config

I use Fish, but it can be applied to other shells.

```bash
switch (uname)
    case Darwin
        # backward compatibility with docker command
        function docker
            podman $argv
        end

        # make it work with tools relying on docker socket
        function lazydocker
            DOCKER_HOST="unix://$(podman machine inspect --format '{{.ConnectionInfo.PodmanSocket.Path}}')" "$HOME/.nix-profile/bin/lazydocker"
        end
end
```
