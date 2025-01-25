+++
date = "2025-01-25"
path = "/posts/2025/01/my-journey-for-fully-switching-from-windows-over-to-linux"
title = "My journey for fully switching from Windows over to Linux"

[taxonomies]
categories = ["infrastructure"]
tags = ["windows", "linux"]
+++

**Note: This is not a post to bash on Windows. It's more of what I need for work and tinkering stuff are better done on
Linux**

## Background

My first OS was Windows. My first-ever automation script was written in Batch script, and I learned tons from tinkering
with the systems with it.

Then I explored Linux during my teens, for around five years I dual-booted Windows and Ubuntu side by side. Love using
both, not willing to lose the other kind of thing.

In college, my laptop was outdated at that point, battery doesn't hold much charge. Apple happened to have discounts for
students, so I got a Macbook. During this time I started learning Python, but I still use Windows for media management
and gaming. I really noticed the difference when programming in Unix-based systems, things are much easier and less
painful.

Then I started working, I'm an avid gamer so I purchased a new PC mostly for gaming and media management (tagging music,
books, etc). Due to my work setup (single monitor, Macbook for work), it means if I want to play games on Windows, I
have to manually switch the monitor input channel and detach USB hub for peripherals attached to a Macbook to the PC
rig. Over time with more work responsibilities, I barely use Windows. This went on for around 5 years.

## The Turning Point

Since I take on management responsibilities, it means my day starts earlier than I was an individual contributor. Which
means I somewhat control my schedule, but it depends on when I have meetings throughout the day.

I drink tea/coffee in the morning, cook every other day during lunch. Previously back when I was an IC, I can not touch
a computer before 1PM, but now I have to stay glued to the computer to attend meetings and respond to bajillion messages
and emails.

I live in a house, due to the layout, my desk and the kitchen island is around 10-meter apart. The problem? I have to move
my Macbook and dock it to my desk. Every single day. It gets old real fast. So I had a bright idea: partition off a 500
GB SSD for gaming on my PC and use it to install Ubuntu on it.

## One Month Later

I really love this setup, then I found out Proton is a thing and gaming on Linux is actually ok. Wheels are turning
in my head and it clicked: I don't need to use Windows for gaming.

To test this idea, I installed Witcher 3 on Ubuntu and gave it a go. Next thing I know is I'm 1/3 finished with the
game. This was really fun. Only problem: The rest of the storage on this PC is tied to Windows. Oops.

## Migration

It would be great if I can move other workloads I still require using Windows fully to Linux. The said workload is music
tagging, which luckily you can use `puddletag` in lieu of `MP3Tag` on Windows. But I still want `foobar2000`-like music
player on Linux, which isn't available. Then I found out someone packaged it as a `snap`. And hey it works!

To make sure I really don't need to use Windows, I migrated everything I need to use Windows for (gaming - done, music
tagging via puddletag - done, foobar2000 - done) and wait for a few weeks.

D-day, I backup data on NTFS to an external drive, remove all partitions, and partition it as follows:

| Drive      | Usage  | Mountpath |
|------------|--------|-----------|
| SSD 256 GB | boot   | /boot/efi |
|            | home   | /home     |
| SSD 500 GB | gaming | /mnt/ssd  |
| HDD 1 TB   | data   | /mnt/hdd  |

Fought a little with Archlinux, because it hates NVIDIA GPU apparently, so had to reflash Ubuntu.

I'm using `Nix` for packages and configurations, in total it took me 1 day to backup + install OS, and the other one for applying the configurations.

## Closing

Overall I'm very happy with this setup, and turns out I can apply better graphics settings for games compared to when I was on Windows.
