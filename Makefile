#
# github.com/MichaelDarr/docker-config
#

BIN := dcon
GO ?= go
VERSION = $(shell cat VERSION)

GOFLAGS := -ldflags "-X github.com/MichaelDarr/docker-config/internal.Version=$(VERSION)"

.PHONY: default
default: $(BIN)

.PHONY: build
build: $(BIN)

.PHONY: $(BIN)
$(BIN): ## build docker-config as dcon
	$(GO) build $(GOFLAGS) -o $(BIN) .

.PHONY: run
run: ## build and run docker-config
	$(GO) run $(GOFLAGS) .
