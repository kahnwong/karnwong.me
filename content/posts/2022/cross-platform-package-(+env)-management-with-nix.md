---
title: Cross-platform package (+env) management with Nix
date: 2022-12-03T19:45:36+07:00
draft: false
ShowToc: false
images:
tags:
  - devops

---

For many years, installing a package on linux means either:

- Compiling a binary from source, then install it. -> I think we know why this didn't catch on for the mass.
- Downloading a compiled binary for your system's architecture and platform. -> This requires you to also move the executable to something like `/usr/local/bin` otherwise it won't be discoverable throughout the system.
- Using system's package manager: `apt`, `apk`, `yum`, `brew`, etc. -> Yay finally something that's easy to use. Phew!

Then the dot-com era happened, and the digital transformation, you name it. This was before cloud, so companies set up their own data centers and have to administer and maintain the servers themselves. And it's not fun if you have to perform the same machine configuration for the whole fleet. This problem was solved by tools like Ansible, Chef, Puppet, etc, to set up a machine's configuration en masse.

So why is this still an issue? Because there are sizable amount of developers / engineers working on a Mac, and deploy to a Unix-based system. This means those machine setup configurations don't work with Mac, and vice versa, because they use different package managers. Technically you can create a separate set of configs for Mac, but this means maintaining two different setup scripts, which aim to do the same thing, just on different platforms. Sooner or later the configuration would get out of sync 😱.

However there are still issues with using system's package manager in linux, namely for some bleeding edge packages, or new package versions that are not yet available in your current linux version, usually involve adding an explicit repository url, adding a keyring, or having to upgrade the system altogether so it can fetch the new repository isn't fun. I've yet to see why we can't install a new package version when its functionalities have almost nothing to do with system version. Otherwise system's package manager is fast and reliable.

Meanwhile there's `brew` for darwin. I only use it because it was one of the few options we have as close as an actual package manager. But it's very slow and performing repo update and installing new apps can take forever, especially if you have to check multiple packages.

Not only that, once you have initialized a system, and over time you modified its configurations, how can you be sure that all the changes are populated back to your Ansible script? This would be a manual process prone to errors.

But humans are awesome, so the best and brightest came up with `nix`, as in `*nix`, a cross-platform package manager that works on both Unix and darwin (a platform name of Mac OS).

With nix, you can utilize `home-manager` to populate the packages/configs on your system, in which it would symlink your configs to nix home, to be symlinked via `home-manager switch` to actual destination, with `read only` file permission. This means if you use nix to initialize `~/.ssh/config` and you want to change it by hand, it would throw "this file has read-only permission" error, this way the only way to update the configs is through nix.

Also with the full system configurations, it takes 6 seconds flat to apply the delta diff 😎.

Some interesting snippets from `home.nix`:

## Initialize dotfiles

```nix
home.file.".ssh/config".source = ./dotfiles/.ssh/config;
```

## Git config

Notice the `delta` block, this automatically populate required configs in `.gitignore`.

```nix
programs.git = {
  # git config --global --edit for raw config content
  enable = true;
  userName = "";
  userEmail = "";

  delta = {
    enable = true;
    options = {
      navigate = true;
      side-by-side = true;
    };
  };

  extraConfig = {
    diff.colorMoved = "default";
    merge.conflictstyle = "diff3";
  };
};
```

## Neovim config

Nix doesn't use `vim-plug` by default, and I found some plugins failed installing via nix, hence the `extraConfig` block for installing vim plugins. Don't forget to run `nvim --headless +PlugInstall +qall`.

```nix
programs.neovim = {
  enable = true;
  viAlias = true;
  vimAlias = true;
  vimdiffAlias = true;

  plugins = with pkgs.vimPlugins; [
    vim-plug
  ];

  extraConfig = ''
    runtime! plug.vim
    call plug#begin()

    "diff
    Plug 'mhinz/vim-signify'

    call plug#end()

    """" config
    set number
  '';
};
```

## Conditional package list for each platform

Maybe you only use linux for servers, so you might not need GUI apps, or some packages are not available on darwin, etc.

```nix
home.packages = with pkgs;
  let
    # Packages to always install.
    common = [
      fish
      starship
    ];

    linux_only = [
      iotop
      ntfs3g
      progress
    ];

    mac_only = [
      mpv
    ];
  in
  common ++ (if stdenv.isLinux then linux_only else mac_only);
```

Head over to the [repo](https://github.com/kahnwong/nix) for my full setup. And remember: great artists _steal_.
