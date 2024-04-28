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
lint: $(BIN)/staticcheck ## Run linters
	$(BIN)/staticcheck ./...

.PHONY: deploy ## Deploy directly to fly.io
deploy:
	fly deploy

$(BIN)/staticcheck: $(BIN) Makefile
	@mkdir -p $(@D)
	GOBIN="$(abspath $(@D))" go install honnef.co/go/tools/cmd/staticcheck@latest
