# See https://tech.davis-hansson.com/p/make/
SHELL := bash
.DELETE_ON_ERROR:
.SHELLFLAGS := -eu -o pipefail -c
.DEFAULT_GOAL := all
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules
MAKEFLAGS += --no-print-directory
BIN := .tmp/bin
export PATH := $(BIN):$(PATH)
export GOBIN := $(abspath $(BIN))

.PHONY: help
help: ## Describe useful make targets
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "%-15s %s\n", $$1, $$2}'

.PHONY: all
all: test lint ## Generate and run all tests and lint (default)

.PHONY: test ## Run all unit tests
test:
	go test -race ./...

.PHONY: lint-go
lint: $(BIN)/staticcheck
	$(BIN)/staticcheck ./...

.PHONY: deploy ## Deploy directly to fly.io
deploy:
	fly deploy

$(BIN):
	@mkdir -p $(BIN)

$(BIN)/staticcheck: $(BIN) Makefile
	go install honnef.co/go/tools/cmd/staticcheck@latest
