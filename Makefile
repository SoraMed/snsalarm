# Go parameters
PROJECT_NAME := $(shell echo $${PWD\#\#*/})
PKG_LIST := $(shell go list ./... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)

all: build

lint: ## Run lint
	@golint -set_exit_status ${PKG_LIST}

vet: ## Run go vet
	@go vet ./...

check: ## Run gosimple and staticcheck
	@gosimple && staticcheck

test: ## Run unittests
	@go test -short ${PKG_LIST}

build: ## Build the binary file
	@GOOS=linux go build
	@zip snsalarm.zip snsalarm

clean: ## Remove previous build
	@go clean ./...

upgrade: ## Get latest libs
	@go get -u

watch:
	@echo Watching for changes...
	@fswatch -or . -e ".*" -i "\\.go$$" | xargs -n1 -I{} make all tags

tags:
	@gotags -R *.go domain controller lib route vendor > tags

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: all lint test race msan dep clean upgrade build watch tags help
