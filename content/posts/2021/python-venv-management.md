---
title: "Python venv management"
date: 2021-07-02T14:59:12+07:00
draft: false
ShowToc: false
images:
tags:
  - python

---

**Updated 2023-09-09**: I now revert back to plain `requirements.txt`, since pipenv is very poor at resolving large dependencies list. Poetry still remains funky so that's off the table as well.

**Updated 2024-01-20**: I've been using poetry in production for a few months now and so far it works great! Looks like they ironed out a lot of rough edges. Although I would still recommend people to use pre-commit to generate plain `requirements.txt`, this is so that in ci or dockerfile you don't have to install poetry and set in-project venv every time.

---

When you create a project in python, you should create `requirements.txt` to specify dependencies, so other people can have the same environment when using your project.

However, if you don’t specify module versions in `requirements.txt`, you could end up with people using the wrong module version, where some APIs can be deprecated or have different behaviors than older versions.

Another issue is that maybe you’re working on a few python projects, each uses different python versions (eg. projectA uses python3.6, projectB uses python3.9, etc).

Enter `pyenv` and `poetry` (`pipenv` is considered outdated now), where you can easily switch python versions, and have different environment (with python version locking) for projects you’re working on.

## Setup pyenv

Follow instructions [here](https://github.com/pyenv/pyenv). For windows, use [this](https://github.com/pyenv-win/pyenv-win).

### Useful commands

```bash
# list available python versions
pyenv install --list

# install specific version
pyenv install 3.8.0

# list installed versions
pyenv versions

# activate new env
pyenv shell 3.8.0 # support multiple version

# config venv
pyenv virtualenv 3.8.0 my-data-project

# set env per folder/project
pyenv local my-data-project
```

## Setup [pipenv](https://github.com/pypa/pipenv)

**2024-01-20: See below for poetry, this only stays for reference only**

Notes: make sure `pyenv` is installed, and remove `anaconda / miconda` & `python3` installed via official installer from your system. Then run:

```bash
$ pip install pipenv

# run this command every time pip installs a .exe
$ pyenv rehash
```

### pipenv workflow

```bash
pipenv --python 3.7

# install a specific module
pipenv install jupyterlab==2.2.9

# install from existing requirements.txt or from Pipfile definition
pipenv install

# remove venv
pipenv --rm

# running inside venv
pipenv run jupyter lab
pipenv run python main.py # is equivalent to `pipenv shell && python3 main.py`
```

## Windows only

**2024-01-20: See below for poetry, this only stays for reference only**

```bash
pyenv install 3.7.7 # see Pipfile for required python version
pyenv local 3.7.7 # IMPORTANT. global / shell doesn't work with pipenv
pyenv rehash
pip install pipenv # done once per pyenv python version
pyenv rehash
pipenv --python 3.7
pipenv install
pipenv run python tokenization_sandbox.py
```

## Setup [poetry](https://python-poetry.org/docs/#installation)

```bash
pipx install poetry
```

### Useful commands

```bash
# init project
poetry init

# add deps
poetry add $package

# add dev dependencies
poetry add $package --group dev

# activate venv
poetry shell

# specify python version
poetry env use 3.11 # normally it picks up global python via pyenv
```

## Notes

* On linux/mac, do not use system python. OS updates would mean python version upgrade, in turn making all your installed modules gone. Use python installed via pyenv instead.
* On windows, start fresh with pyenv.
* Do not use anaconda distribution. It does too much background magic that can make things harder to manage environment properly. In addition, venv definition from anaconda is often doesn’t work cross-platform (eg. venv def from windows wouldn’t work on mac due to different wheel binary versions).
* Always create venv via pipenv per each project. Although you can have a playground venv via pyenv, so you can shell into it and do a quick analysis / scripting on an adhoc basis.
* ~~I heard good things about [poetry](https://github.com/python-poetry/poetry) but it doesn’t integrate with `pyenv` natively. It would work if you use it to publish python modules, since it simplifies a lot of processes.~~
  * ~~poetry also picks up the wrong python version from pyenv. And if you sync python version via pyenv, it has to be the same python version across all OSes, including minor version. pipenv doesn’t have this restriction, and it also picks up the correct python version from pyenv by default (via `pipenv --python 3.8`).~~
