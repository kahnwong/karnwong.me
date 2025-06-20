+++
title = 'Using Nix for CI, but with a twist'
date = '2025-06-20'
path = '/posts/2025/06/using-nix-for-ci-but-with-a-twist'

[taxonomies]
categories = [ 'ci-cd',]
tags = []
+++

CI (continuous integration) is kind of tricky in a sense that each CI vendor has its own implementations, often glued together via yaml configurations. If you have a CI to build a binary, unless you use bash to 1) install dependencies and b) build the artifacts, most likely you would end up with different implementation for each CI platform.

The problem with using plain bash script is that it's a lot of glue code (setting PATH variable for one), not to mention having to make sure that dependencies are pinned to a specific version (because things are going to break with upgrades, I can assure you that).

This essentially means that your CI pipelines would be very well optimized if you use vendor-specific building blocks (for installing dependencies, building binaries with cache, etc), but it's not portable. Bash scripts are portable but very brittle in a sense that it's harder to pin dependencies.

But Nix exists, so why not. Nix is a build system, but you can use it to install packages. Take [gomod2nix](https://github.com/nix-community/gomod2nix) for example, basically it stays within your golang project so you can use nix to build a golang binary. It's very pleasant to use, but that's a lot of boilerplate code, and it's still tricky to figure out how to do cross-compilation build.

But then [flox](https://flox.dev/docs/install-flox/) exists. Think of virtual environment, but behind the scenes it's nix installing tools and packages. This means you can offload dependencies installation to flox/nix, and continue using bash for the build process.

In `.flox/env/manifest.toml`, you would specify packages like this:

```toml
[install]
go = { pkg-path = "go", version = "^1.23.3" }
air.pkg-path = "air"
```

And once you activate the flox shell, you can continue doing your own thing.

In CI, this means that as long as you can install flox/nix, and set it to utilize cache, you are pretty much good to go. Although you probably have to use vendor-specific configurations for uploading build artifacts, but it's a small price to pay for not having to reimplement everything for each CI system.
