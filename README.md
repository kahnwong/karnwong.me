# blog

## Initial setup
```bash
git clone $REPO
git submodule add git@github.com:kahnwong/hugo-theme-hello-friend-ng.git themes/
git submodule update --init --recursive
```

## Update theme
```bash
git pull --recurse-submodules
```

## Create a new post
```bash
./create-post.sh -n $TITLE -t $TAG
```
