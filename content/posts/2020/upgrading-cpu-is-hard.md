---
title: CPU upgrade is a breeze, only if you know how
date: 2020-12-20T19:12:38.000Z
draft: false
ShowToc: false
images:
tags:
  - life
---

I want to update my pc config to use a more powerful cpu (from Ryzen 1300x --> Ryzen 3600). Watched a few videos and it didn't look complicated enough, so I order the new CPU and plan to install it myself. Things didn't go exactly to plan. To sum up relevant things I didn't know from watching "how to install AM4 CPU" videos:

- on some motherboard, you need to "support" the backplate from under the motherboard so the socket would protrude far up enough to "lock" the CPU fan.
- even if ab350 says it supports Ryzen 3rd generation CPUs after flashing the supported BIOS version, it  is *not* a given that it will work --> I end up upgrading my motherboard to b450
- if you boot Windows after changing CPU / GPU, it won't boot --> have to reset CMOS

Of course, I didn't figure all this out myself. I took the pc to a local repair shop to install the CPU fan but it didn't POST. We thought it's the from the CPU pins I accidentally bent and maybe I screw it up when I straighten it back.

But it's still within 7 days of purchase, so I took it to the retailer for a claim, only to find out that it works. The tech ask me which motherboard I have, only then it hit them that I need to upgrade the motherboard.

Then I took it back to a repair shop to install the new motherboard + CPU, but it still doesn't POST. Figured it out later that you need to reset CMOS.

The "how to install AM4 CPU" didn't lie or anything, it's just that all the above points are "basic knowledge" for everyone who works on pc hardware should know. It's the same as "try a newly installed CLI from a new terminal session" <– because your shell pre-index executables when it inits a session.

Takeaway: don't underestimate domain knowledge basics.
