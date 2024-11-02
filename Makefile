# Binary name
BINARY_NAME=grimoire

# Go related variables
GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin
GOFILES=$(wildcard *.go)

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

.PHONY: all build clean test coverage run fmt lint help


build: ## Build the binary
	@echo "Building..."
	go build -o $(GOBIN)/$(BINARY_NAME) ./cmd/grimoire


run: build ## Build and run the binary
	$(GOBIN)/$(BINARY_NAME)


clean: ## Clean build files
	@echo "Cleaning..."
	go clean
	rm -rf $(GOBIN)

## Run tests
test:
	go test ./... -v


coverage: ## Run tests with coverage
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out


fmt: ## Format code
	go fmt ./...

lint: ## Lint code
	go vet ./...


help: ## Display this help screen
	@echo "Usage:"
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'