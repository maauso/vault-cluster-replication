
GOARCH ?= amd64
GOOS ?= linux
DOCKER_IMAGE ?= vault-cluster-replication



# Go application name
APP_NAME = vault-cluster-replication

# Go source files
SRC = $(wildcard *.go)

# Binary name for test coverage report
COVERAGE_BINARY = coverage.out

# Golangci-lint Docker image
LINT_IMAGE = golangci/golangci-lint:v1.53-alpine

# Go Docker image for building
GO_IMAGE = golang:1.20-bookworm

.PHONY: all test lint build build-docker-image

all: test lint build

test:
	@go test -v ./...

coverage:
	@go test -coverprofile=$(COVERAGE_BINARY) ./...

lint:
	@docker run --rm -v $(PWD):/app -w /app $(LINT_IMAGE) golangci-lint run

clean:
	@rm -f $(COVERAGE_BINARY)


build-docker-image:
	@docker build --build-arg GOARCH=$(GOARCH) -f ./Dockerfile -t $(DOCKER_IMAGE) .

build:
	@GOARCH=$(GOARCH) GOOS=$(GOOS) go build -o dist/$(GOARCH)/bin/vault-cluster-replication ./cmd/server/main.go

.PHONY: clean
