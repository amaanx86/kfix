PHONY: build test lint install clean release help

BINARY_NAME=kfix
VERSION?=$(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
LDFLAGS=-ldflags "-X github.com/amaanx86/kfix/cmd.version=$(VERSION)"

help:
	@echo "kfix - Kubernetes YAML Formatter"
	@echo ""
	@echo "Available targets:"
	@echo "  build         - Build the binary"
	@echo "  test          - Run tests"
	@echo "  test-verbose  - Run tests with verbose output"
	@echo "  lint          - Run golangci-lint"
	@echo "  install       - Install the binary"
	@echo "  clean         - Remove the binary"
	@echo "  release       - Build binaries for all platforms"
	@echo "  help          - Show this help message"

build:
	@echo "Building kfix..."
	go build $(LDFLAGS) -o $(BINARY_NAME) .

test:
	@echo "Running tests..."
	go test -v ./...

test-verbose:
	@echo "Running tests with verbose output and coverage..."
	go test -v -cover ./...

coverage:
	@echo "Generating coverage report..."
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

lint:
	@echo "Running linter..."
	golangci-lint run

install: build
	@echo "Installing kfix..."
	@if [ "$(shell uname -s)" = "Darwin" ] || [ "$(shell uname -s)" = "Linux" ]; then \
		echo "Installing to /usr/local/bin..."; \
		sudo mv $(BINARY_NAME) /usr/local/bin/$(BINARY_NAME); \
	else \
		echo "Unsupported OS for this install method. Using 'go install' instead."; \
		go install $(LDFLAGS); \
	fi

clean:
	@echo "Cleaning up..."
	rm -f $(BINARY_NAME)
	rm -f coverage.out coverage.html

release:
	@echo "Building releases for all platforms..."
	mkdir -p dist
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o dist/$(BINARY_NAME)-linux-amd64 .
	GOOS=linux GOARCH=arm64 go build $(LDFLAGS) -o dist/$(BINARY_NAME)-linux-arm64 .
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o dist/$(BINARY_NAME)-darwin-amd64 .
	GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o dist/$(BINARY_NAME)-darwin-arm64 .
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o dist/$(BINARY_NAME)-windows-amd64.exe .
	@echo "Releases built in dist/"

fmt:
	@echo "Formatting code..."
	go fmt ./...

vet:
	@echo "Vetting code..."
	go vet ./...

check: fmt vet lint test
	@echo "All checks passed!"

.DEFAULT_GOAL := help