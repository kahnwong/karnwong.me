#!/bin/bash

while getopts n:t: flag
do
    case "${flag}" in
        n) title=${OPTARG};;
        t) tag=${OPTARG};;
    esac
done

slug=$(echo "$title" | awk '{print tolower($0)}')
slug=$(echo "$slug" | tr " " -)

date=$(date "+%Y-%m-%dT%H:%M:%S%Z:00")
date_str=$(date "+%Y-%m-%d")

echo "---
title: "$title"
date: "$date"
draft: false
ShowToc: false
images:
tags:
  - "$tag"
---" > "content/posts/$date_str-$slug.md"
