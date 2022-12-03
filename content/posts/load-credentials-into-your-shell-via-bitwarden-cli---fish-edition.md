---
title: Load credentials into your shell via Bitwarden CLI - Fish edition
date: 2022-11-29T05:59:29+07:00
draft: false
ShowToc: false
images:
tags:
  - devops
---

Recently I work with GitHub CLI a lot, and having to constantly fire up Bitwarden app to retrieve `GITHUB_TOKEN` gets old real fast...

I was thinking of storing it in a gist in a password manager, luckily someone had the same idea and [implemented it](https://blog.gruntwork.io/how-to-securely-store-secrets-in-bitwarden-cli-and-load-them-into-your-zsh-shell-when-needed-f12d4d040df). The only issue is that I use fish shell. But we live in a world where there are many ways to interact with the shell, so it follows that you can translate zsh syntax to fish syntax.

For the original snippet in zsh, translated as fish:

```fish
function unlock_bw_if_locked
    if test -z BW_SESSION
        echo 'bw locked - unlocking into a new session'
        export BW_SESSION="$(bw unlock --raw)"
    end
end

function load_github
  unlock_bw_if_locked

  set -l github_pat_id $BITWARDEN_GIST_ID
  set -l github_token
  set -l github_token "$(bw get notes $github_pat_id)"
  export GITHUB_OAUTH_TOKEN="$github_token"
  export GITHUB_TOKEN="$github_token"
  export GIT_TOKEN="$github_token"
end
```
