+++
title = 'Simplify SSH (and HTTP & Postgres) access management with Warpgate'
date = '2025-06-28'
path = '/posts/2025/06/simplify-ssh-and-postgres-http-access-management-with-warpgate'

[taxonomies]
categories = [ 'devops' ]
tags = ['security']
+++

## Background

If you have to work with virtual machines (VMs), you know the drill: disable password authentication and only allow private key authentication. If you want to sleep better at night, make the VMs only accessible through a private network to reduce the attack surface.

What about at scale? What if you are not the only person who has to access those VMs, because you are working in a team. Then things get interesting, because allowing developers direct access to VMs means a lot of security implications, and it's not uncommon for developers in a hurry to share their private key with a team member - and we all know that the key wouldn't be rotated anytime in the near future.

## How it's been done

In practice, this mostly translates to a bastion host, acting as a bridge between developer's machines and target machines. First you use the bastion host to forward the SSH port, then you use it to access the target machine. This would likely involve the bastion host admin to provision a unique private key for each developer, add its public keys to the bastion host. But this is only the bastion host we are talking about, then there are those target machines where ideally the same public keys should be added to authorized keys as well. You see where this is going. If someone leaves the org, yup, better make sure that you revoke their public key from all machines.

One thing to keep in mind is you'll have to add some guardrails to disable shell access, since some developers, in a hurry, against their better judgment, might add a few stowaway cronjobs critical to production to the bastion host 🙃.

### The hard part that we all dread, but they need to be done

Did I mention audit logs? Have fun combing through the logs. And if you find provisioning private keys tedious, I'm with you on that. Then you have to come up with clever tricks to make sure only certain developer's group has to access to particular VMs, while disable access to other VMs. RBAC strikes back.

## Then I found Warpgate

I'll admit at first it was tricky to assess what warpgate was trying to solve, but upon reading their [documentation](https://warpgate.null.page/) it paints a clearer picture - to use RBAC to control access to targets.

This comprises three layers:

1. User Identity - can be native to warpgate or via SSO
2. Targets - can be SSH, HTTP, MySQL or Postgres

> I find this is very neat since I don't need separate systems for managing non-SSH access.

3. RBAC, which users/groups can access which targets, and which authentication mechanism to use

> You can use password, private key or web-based approval workflow. They can be specified separately per each target type as well.

The setup works like this: warpgate generates its own private/public key on initialization. You only have to add its public key to target machines, this means for all your VMs fleet, only one public key is required to be added in order for warpgate to be able to access it. When you SSH into your target machine, your host would be `warpgate's acme` and username is `$WARPGATE_USERNAME:$WARPGATE_TARGET_NAME`. You can supply a password or private key, or use the web-based approval flow for authentication, depending on how you set it up, and you'll be forwarded to the target machine.

It is similar for HTTP target, but the cool thing is you can bind it to a domain as well, which makes it great for internal apps. As for databases, it is similar to ssh targets, existing clients should work without requiring extra configurations.

And per ops tradition, warpgate has a terraform provider, and you should definitely use terraform to manage RBAC, because you'll breeze through the audit.

But let's not forget one of those times when you need to grant a temporary access, warpgate has you covered as well.

---

The best part is that warpgate is free, as in beer 🍺.
