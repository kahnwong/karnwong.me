# blog

## Initial setup
```
git clone $REPO
git submodule update --init --recursive
```

## Update theme
```
git submodule foreach git pull origin master
```

## Create a new post
```
./create-post.sh -n hello -t test
```
