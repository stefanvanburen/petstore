# See https://tech.davis-hansson.com/p/make/
SHELL := bash
.DELETE_ON_ERROR:
.SHELLFLAGS := -eu -o pipefail -c
.DEFAULT_GOAL := all
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules
MAKEFLAGS += --no-print-directory
BIN := .tmp/bin

.PHONY: help
help: ## Describe useful make targets
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "%-15s %s\n", $$1, $$2}'

.PHONY: all
all: test lint ## Generate and run all tests and lint (default)

.PHONY: test
test: ## Run all unit tests
	go test -race ./...

.PHONY: lint
lint:  ## Run linters
	go tool honnef.co/go/tools/cmd/staticcheck ./...

.PHONY: deploy
deploy: ## Deploy directly to fly.io
	fly deploy
