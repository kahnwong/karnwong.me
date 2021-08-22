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