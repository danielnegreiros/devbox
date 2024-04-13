# Makefile for pre-commit checks in a Go project

# Variables
COVERAGE_FILE := coverage.out
GOLANGCI_LINT := $(GOPATH)/bin/golangci-lint

# Set default goal
.DEFAULT_GOAL := lint

# Tidying
tidy:
	go mod tidy

# Linting
lint: $(GOLANGCI_LINT)
	$(GOLANGCI_LINT) run
	go vet ./...

# Run tests
test:
	go test -v ./...

# Generate coverage report
coverage:
	go test -coverprofile=$(COVERAGE_FILE) ./...
	go tool cover -func=$(COVERAGE_FILE)

# Format code
format:
	gofmt -s -w .

# Profile code
profile:
	go test -bench=. ./...

build:
	go build -o devbox cmd/proxmox/main.go
	sudo mv devbox /usr/local/bin/

all: tidy lint test coverage format profile build

# Ensure golangci-lint is installed
$(GOLANGCI_LINT):
	@if [ ! -f $(GOLANGCI_LINT) ]; then \
		echo "Installing golangci-lint..."; \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.57.1; \
	fi
