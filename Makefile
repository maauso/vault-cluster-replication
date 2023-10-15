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

.PHONY: all test lint build

all: test lint build

test:
	@go test -v ./...

coverage:
	@go test -coverprofile=$(COVERAGE_BINARY) ./...

lint:
	@docker run --rm -v $(PWD):/app -w /app $(LINT_IMAGE) golangci-lint run

clean:
	@rm -f $(COVERAGE_BINARY)

.PHONY: clean
