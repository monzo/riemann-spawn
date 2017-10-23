VERSION=0.1
APP_NAME=transactcharlie/riemann-spawn
# HELP
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help clean

help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help


# DOCKER TASKS
dockerBuild: ## Build the container using an intermediary go build env
	docker build \
		-t $(APP_NAME):$(VERSION) \
		-t $(APP_NAME) \
		.

build: dockerBuild
