# blog

## Initial setup

```bash
git clone $REPO
git submodule add git@github.com:kahnwong/hugo-PaperMod.git themes/PaperMod
git submodule update --init --recursive
```

## Update theme

```bash
git submodule foreach git pull origin master
```

## Create a new post

```bash
./create-post.sh -n $TITLE -t $TAG

# with images
./create-post.sh -n $TITLE -t $TAG -i true
```
