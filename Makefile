start:
	zola serve
build:
	zola build
theme-init:
	git submodule update --init --recursive
theme-update:
	git submodule foreach git pull origin main
resume-build:
	cd resume && yarn build && cp resume.html ../static/
