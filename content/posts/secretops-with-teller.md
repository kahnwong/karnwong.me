---
title: SecretOps with teller
date: 2023-03-19T00:57:14+07:00
draft: false
ShowToc: false
images:
tags:
  - devops
  - vault
  - secretops
---

Raise your hands if you normally have to send `.env` files to your team members so they can start a project in dev environment.

While there is nothing wrong with this approach, it could introduce a lot of security risks, namely sharing secrets via plaintext protocol. Sure, you can share them on Slack, Discord, etc. But unless it's encrypted with your own keys, it could be leaked if the said communication platform were to be breached.

You could uses a one-time pastebin to share the secrets. This would be a step up, but you still have to manually send each team member the secrets.

Some might say there are services like <https://keybase.io> where you can send messages encrypted with PGP key. Yay for security, but it has the same issue as above.

You use cloud, right? So most if not all cloud providers provide their own secrets management solution, so let's use that. But then developers have to somehow fetch the secrets, and set it as environment variable so they can start a project. This would work until you update the dev secrets, and some developers would eventually forget to pull a new set of secrets before starting a project. This could result in half a day of wack-a-mole hunt for why it's not working on their machine, but works on their friends'. Or you don't, if you fetch secrets during a project initialization. But the [twelve-factor-app](https://12factor.net/) would like a word.

I stumbled upon [teller](https://github.com/tellerops/teller) accidentally, but it seems to be able to do what I want: abstracting away secrets fetching and populate it as environment variable in one command.

Setting up is as easy as: `brew tap spectralops/tap && brew install teller`.

To populate secrets as environment variables: `eval "$(teller sh)"`

I have a [repo](https://github.com/kahnwong/teller-playground) here you can tinker with. Please let me know if you find some quirks, since I plan to use this with my team next year 🎊.
