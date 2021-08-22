title="Title here"

slug=$(echo "$title" | awk '{print tolower($0)}')
slug=$(echo "$slug" | tr " " -)

date=$(date "+%Y-%M-%dT%H:%M:%S%Z:00")
tag="sample-tag"

echo "---
title: "$title"
date: "$date"
draft: false
toc: false
images:
tags:
  - "$tag"
---" > "content/posts/$slug.md"
