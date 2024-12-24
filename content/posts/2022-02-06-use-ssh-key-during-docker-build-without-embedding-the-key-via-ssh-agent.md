+++
title = "Use SSH key during Docker Build without embedding the key via ssh-agent"
date = "2022-02-06"
path = "/posts/2022/02/use-ssh-key-during-docker-build-without-embedding-the-key-via-ssh-agent"

[taxonomies]
categories = ["ci-cd",]
tags = [  "docker", "github",]

+++

Imagine working in a company, and they have a super cool internal module! The module works great, except that it is a private module, which means you need to install it by cloning the source repo and install it from source.

That shouldn't be an issue if you work on your local machine. But for production usually this means you somehow need to bundle this awesome module into your docker image. You go create a Dockerfile and there's one little problem: it couldn't clone the module repo because it doesn't have the required SSH key that can access the repo.

A very simple solution would just be bundling the SSH key into the docker image itself. This works great, until security comes knocking because: anyone who can access the image can also access the source repo! And we wouldn't want that.

So what's the solution? Luckily there's a thing called ssh-agent forwarding. Think of it as passing an SSH key into docker during build, then poof! It's gone.

## Instructions

1. Set up SSH key on your host machine. This essentially means you should be able to perform `git clone $REPO` on the module repo.
2. In your Dockerfile, add something like the following:

```dockerfile
# Authorize SSH Host
RUN mkdir -p -m 700 /root/.ssh && \
    touch -m 600 /root/.ssh/known_hosts && \
    ssh-keyscan github.com > /root/.ssh/known_hosts

RUN --mount=type=ssh,id=github $SOME_COMMAND_THAT_NEEDS_SSH_KEY
```

3. Build docker image with `docker build --ssh github=$SSH_PRIVATE_KEY_PATH -t $IMAGE_NAME .`

## GitHub Actions

In GitHub actions, #1 is translated as:

```yaml
- name: Setup SSH Keys and known_hosts
  run: |
    mkdir -p ~/.ssh
    ssh-keyscan github.com >> ~/.ssh/known_hosts
    ssh-agent -a $SSH_AUTH_SOCK > /dev/null
    ssh-add - <<< "${{ secrets.SSH_PRIVATE_KEY }}"
  env:
    SSH_AUTH_SOCK: /tmp/ssh_agent.sock

- name: Build, tag, and push image to Amazon ECR
  uses: docker/build-push-action@v2
  env:
    SSH_AUTH_SOCK: /tmp/ssh_agent.sock
  with:
    context: .
    push: true
    tags: ${{ steps.meta.outputs.tags }}
    cache-from: type=gha
    cache-to: type=gha,mode=max
    ssh: |
      github=${{ env.SSH_AUTH_SOCK }}
```

Voila 🎉
