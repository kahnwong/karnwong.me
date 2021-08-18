export title=""
export slug=""
export date=""
export tag=""

echo "---
title: "$title"
date: "$date"
draft: false
toc: false
images:
tags:
  - "$tag"
---" > "content/posts/$slug.md"
