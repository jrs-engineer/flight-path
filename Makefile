VERSION := $(shell git describe --tags --always)
GITREV := $(shell git rev-parse --short HEAD)
GITBRANCH := $(shell git rev-parse --abbrev-ref HEAD)
DATE := $(shell LANG=US date +"%a, %d %b %Y %X %z")

GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/dist
GOENVVARS := GOBIN=$(GOBIN) CGO_ENABLED=0
GOBINARY := flight-parser
GOCMD := $(GOBASE)/cmd/

LDFLAGS += -X 'github.com/C001-developer/flight_path.Version=$(VERSION)'
LDFLAGS += -X 'github.com/C001-developer/flight_path.GitRev=$(GITREV)'
LDFLAGS += -X 'github.com/C001-developer/flight_path.GitBranch=$(GITBRANCH)'
LDFLAGS += -X 'github.com/C001-developer/flight_path.BuildDate=$(DATE)'


# Building the docker image and the binary
build: ## Builds the binary locally into ./dist
	$(GOENVVARS) go build -ldflags "all=$(LDFLAGS)" -o $(GOBIN)/$(GOBINARY) $(GOCMD)
.PHONY: build

# Linting, Teseting, Benchmarking
golangci_lint_cmd=github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.2

install-linter:
	@echo "--> Installing linter"
	@go install $(golangci_lint_cmd)

lint:
	@echo "--> Running linter"
	@ $$(go env GOPATH)/bin/golangci-lint run --timeout=10m
.PHONY:	lint install-linter

test: ## Runs the tests
	go test ./... --timeout=10m
.PHONY: test

run: build
	./dist/flight-parser
.PHONY: run