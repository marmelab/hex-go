.PHONY: default install start test stop clean

.DEFAULT_GOAL := help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

UID = $(shell id -u)
GID = $(shell id -g)
BIN = docker-compose run \
    --rm hex-go

start: ## Start project
	@echo "Start the project"
	$(BIN) go run /go/src/hex/main.go

stop: ## Stop project
	@echo "Stop the project"
	docker-compose down

test: ## Launch the project's tests
	@echo "Launch the tests"
	$(BIN) go test -v ./src/hex/...
	$(BIN) go test -v ./src/api/...

clean: ## Clean the project
	@echo "Cleaning project"
	$(BIN) rm -Rf pkg bin

install: ## Install project's dependencies
	@echo "Install project deps"
	docker-compose build
	$(BIN) go get ./src/hex/...
