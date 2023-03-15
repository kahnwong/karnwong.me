#!/bin/bash

while getopts n:t:i: flag; do
	case "${flag}" in
	n) title=${OPTARG} ;;
	t) tag=${OPTARG} ;;
	i) contains_image=${OPTARG} ;;
	esac
done

slug=$(echo "$title" | awk '{print tolower($0)}')
slug=$(echo "$slug" | tr " " -)

date=$(date "+%Y-%m-%dT%H:%M:%S%Z:00")
date_str=$(date "+%Y-%m-%d")

if [[ -n "$contains_image" ]]; then # create a post with images
	mkdir -p "content/posts/$slug"
	mkdir -p "content/posts/$slug/images"

	echo "---
title: "$title"
date: "$date"
draft: false
ShowToc: false
images:
tags:
  - "$tag"
---" >"content/posts/$slug/index.md"
else
	echo "---
title: "$title"
date: "$date"
draft: false
ShowToc: false
images:
tags:
  - "$tag"
---" >"content/posts/$slug.md"
fi
