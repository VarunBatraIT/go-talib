ALL_PKGS := $(shell go list ./... | grep -v /vendor)

GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOLANGCI-LINT=golangci-lint

.PHONY: setup_repo run run_heap clean_all test lint

# DON'T EDIT BELOW

all: test_ci

clean_all:
	$(GOCMD) clean ./...
	rm -rf ./release/

test_ci:
	$(GOCMD) test -cover -race -v $(ALL_PKGS)

test_all:
	$(GOCMD) test -race -v $(ALL_PKGS)

lint:
	$(GOLANGCI-LINT) run -verbose

setup_repo:
	$(shell pwd)/scripts/repo_setup/run.sh
