+++
title = "Thoughts on summarization service system design"
date = "2024-06-09"
path = "/posts/2024/06/thoughts-on-summarization-service-system-design"

[taxonomies]
categories = ["software-engineering",]
tags = [  "system-design"]

+++

For a summarization task, there should be an input, in which it's reduced to a handful of paragraphs. This input is in text format.
You don't necessarily start from a text format though, since the source content can be audio or video files. But this means at the end, the source input has to be converted into a text format, and this involves a `transcription` task.

Transcription means taking an audio, then convert it to text. Luckily these days there are APIs you can use to achieve this. Depending on each API provider, but it's safe to assume most would support `WAVE` or `FLAC` encoding.

## Audio formats

Audio formats can be categorized into two main groups: lossless and lossy. Lossless means all information are preserved, and can be converted between any lossless format without losing data. For example, going from `WAVE -> FLAC` should allow you to convert the resulting `FLAC -> WAVE` without losing quality.

However, when you convert `MP3 -> WAVE`, this is a lossy to lossless format, which means the resulting `WAVE` file would contain as much information as the MP3 file.

## Why does audio format matter

If you read [GCP's Cloud Speech-to-Text docs](https://cloud.google.com/speech-to-text/docs/encoding#lossless_compression), it would tell you that using `FLAC` is better if you are conscious about storage.
This is good and all, but the caveats (not mentioned in the docs) is that, if you convert a lossy format (which is very common in videos) to FLAC, it would take a lot longer than converting it to WAVE. This is because WAVE is an `uncompressed` format, which means there's less conversion overhead, seeing FLAC utilizes a heavy compression.

We are talking about `48 hours` conversion time vs `a few minutes`. Compute is cheap, but storage is cheaper.

## What about calling patterns

Both transcription and summarization can take quite a while (longer than a few minutes). It does not make sense to throw the entire logic into a function and have your api return it as-is, since it would leave a connection open, and if your workloads are interrupted, all progress are lost.
Which is why an event-based approach is a better fit for this scenario. Initial api call would trigger an event, and serverless functions (one for each transcription and summarization) run these tasks, when completed a status is updated into the database. Frontend can poll the backend every few seconds to fetch the status of these long-running operations.

## What if you are lazy and just want to only create a single backend?

Sure, it can work. As per anything in this universe: if you try hard enough, it will somehow work. But good luck explaning to your users why it takes forever to see a response on the web UI.
