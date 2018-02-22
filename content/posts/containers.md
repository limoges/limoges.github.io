---
title: "Virtualization technology and containers"
date: 2015-03-02
description: "A brief history of virtualization technologies"
lastmod: 2017-02-04
draft: false
tags: ["docker"]
categories: ["history", "software"]
---
The world of software has been hearing about containers, docker, and docker for 
the past few years.

Probably to the most promoted and **hyped** technology since, maybe, the iPhone, one
should ask himself, what is Docker?

*(2017: Okay. Microservices, k8s - that's kubernetes for the rest of us, are also
strong contenders.)*

Docker is a somewhat recent technology, or tool, to allow *building a workload* and
*running that workload* on a host of different machines.

This technology is key in the era of **everything in the cloud** that we 
live in, where hardware has become a commodity into which software is inserted
in order to run, similar to washing clothes at a laundromat.

## A little bit of history

A few years ago, the commonly accepted methodology for running software in the
cloud involved running many computers on a single computer. 

It involved having a host create a virtual machine (VM) onto which an operating
system (OS) would be booted in order to run your workload. This methodology was
the focus of so-called "virtualization". This virtualization allowed people who
owned hardware to rent it to people who would pay for virtual machines of the
size they desired, depending on their needs.

Many services and tools were built to allow people to jam their application's
requirement into a box which could then be put to run on the virtual hardware.
This is how the cloud computing concept emerged. You didn't need to buy hardware,
you could just pay for what you needed to use, with a certain granularity.

Different vendors had different idea of how they could organize their client's 
software to run onto their hardware, with varying degrees of efficiency and success.
One thing it did though was that it improved hardware utilization across the board,
allowing everyone's [coffee machine] to connect to the internet.

That whole machine virtualization had a cost though. It required an hypervisor
to run on every machine, and a virtual machine to be instantiated for every
workload. Creating, running and destroying these virtual machines took time and 
resources which meant all of these components had a pretty significant overhead.
And that overhead was spent for every single workload.

Turns out, most workload were simple web services which spent a lot of time doing
absolutely nothing. A lot of money was poured into the lucrative virtualization
business to lower this overhead.

## Virtualization using containers

Containers are a virtualization technology. They allow running workload on shared
hardware, but they do so differently. Instead of running an hypervisor to
isolate workloads on different virtual machines, containers use modern kernel
sandboxing facilities isolate processes running on the host.

Getting rid of the hypervisor and virtual machines altogether allows much smaller
resource usage per application. Those applications are set up and torn down much
faster allowing a much higher number of application to run on the same hardware
at anytime - application density. As you might imagine, there are pretty solid
incentives to adopt this kind of technology. Think 10x to 100x of incentives.

### Okay, now what about Docker?

Docker is simply an extension of the Linux Containers (LXC). 
It adds tooling, and provides an API to run processes using LXC technologies. 

This means that with Docker, you can create a portable deployment unit, version
those units (known as containers) and reuse these units.

In essence, Docker is lowering the barrier to using container instead of 
hardware virtualization.

