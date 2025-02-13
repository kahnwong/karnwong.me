+++
date = "2025-02-13"
path = "/posts/2025/02/cross-platform-keyring-like-setup-also-works-without-desktop-environment"
title = "Cross-platform keyring-like setup (also works without desktop environment)"

[taxonomies]
categories = ["infrastructure"]
tags = ["linux", "windows", "darwin", "secrets"]
+++

Sometimes, running an application requires you to supply a password, or something along these lines. If you have a terminal-based workflow, sometimes it can require an authentication token.

These configurations would need to be supplied either at runtime or stored somewhere on your system. We can agree that typing the creds each time isn't fun. But putting those secrets in a file on your filesystem? Attackers love this, so maybe don't make them too happy!

Conventional wisdom is you should store the credentials in a system's keyring. This means you need a desktop environment (headless linux won't work with this). And that if you were to migrate to a new machine, you have to set it up again. Hmmmm.....

I ran into this exact predicament, and I'm happy to share you my workaround to this. It's very simple - using [SOPs](https://github.com/getsops/sops) to encrypt secrets and store it somewhere on your system. Then create a shell function to fetch this secret key from your encrypted file.

## Usage

1. Setup SOPs: <https://github.com/getsops/sops#usage>. I use `age`.
2. Create a secrets file: `sops -i secrets.sops.yaml`
3. Write your secret keys and values
4. Move your `secrets.sops.yaml` to your shell config's path. I use `fish` so it's at `$HOME/.config/fish/secrets/`.
5. Create a fish function to fetch a secret via key

```bash
function get_fish_secret
    ~/.nix-profile/bin/sops -d ~/.config/fish/secrets/secrets.sops.yaml | ~/.nix-profile/bin/yq .$argv
end
```

6. In your shell alias/functions/config, access secrets via `$(get_fish_secret MINIO_ENDPOINT)`

---

This setup also means you can store this with your dotfiles, since it's already encrypted.

Here's to KISS, and don't let anyone tell you otherwise.
