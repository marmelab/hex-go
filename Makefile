.PHONY: default install start test

.DEFAULT_GOAL := help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

UID = $(shell id -u)
GID = $(shell id -g)
BIN = docker-compose run \
    --rm hex-go

install: ## Install project's dependencies
	@echo "Install project deps"
	docker-compose build

start: ## Start project
	@echo "Start the project"
	docker-compose up

test: ## Launch the project's tests
	@echo "Launch the tests"
	$(BIN) go test -v ./src/hex
