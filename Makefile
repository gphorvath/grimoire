# Binary name
BINARY_NAME=grimoire

# Shell variables
USER_SHELL := $(shell echo $$SHELL)
SHELL_NAME := $(shell basename $(USER_SHELL))
COMPLETION_DIR := $(HOME)/.local/share/bash-completion/completions
ZSH_COMPLETION_DIR := $(HOME)/.zsh/completion
FISH_COMPLETION_DIR := $(HOME)/.config/fish/completions

# Go related variables
GO_BASE=$(shell pwd)
GO_BIN=$(GO_BASE)/bin
LOCAL_DIR=~/.grimoire

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

.PHONY: all build clean help completion


build: ## Build the binary
	@echo "Building..."
	@echo "Current directory: $(GO_BASE)"
	@go build -o $(GO_BIN)/$(BINARY_NAME) .

install: build ## Install the binary and prompts
	@echo "Installing Grimoire..."
	mkdir -p $(LOCAL_DIR)
	cp -r prompts $(LOCAL_DIR)
	@echo "[Requires sudo] Copying binary to /usr/local/bin/$(BINARY_NAME)"
	sudo cp $(GO_BIN)/$(BINARY_NAME) /usr/local/bin/$(BINARY_NAME)
	@echo "Done!"
	@echo "Running Grimoire..."
	$(BINARY_NAME)

uninstall: ## Uninstall the binary and prompts
	@echo "Uninstalling..."
	@echo "[Requires sudo] Removing binary from /usr/local/bin/$(BINARY_NAME)"
	sudo rm -f /usr/local/bin/$(BINARY_NAME)
	rm -rf $(LOCAL_DIR)
	@echo "Done!"

completion: install ## Install shell completion
	@echo "Installing completion for $(SHELL_NAME)"
	@case $(SHELL_NAME) in \
		bash) \
			mkdir -p $(COMPLETION_DIR) ; \
			grimoire completion bash > $(COMPLETION_DIR)/grimoire ;; \
		zsh) \
			mkdir -p $(ZSH_COMPLETION_DIR) ; \
			grimoire completion zsh > $(ZSH_COMPLETION_DIR)/_grimoire ;; \
		fish) \
			mkdir -p $(FISH_COMPLETION_DIR) ; \
			grimoire completion fish > $(FISH_COMPLETION_DIR)/grimoire.fish ;; \
		*) \
			echo "Unsupported shell: $(SHELL_NAME)" >&2 ; \
			exit 1 ;; \
	esac
	@echo "Installed completion for $(SHELL_NAME)"
	@echo "Restart your shell or source the completion file to use it"

clean: ## Clean build files
	@echo "Cleaning..."
	go clean
	rm -rf $(GO_BIN)
	rm -rf $(LOCAL_DIR)

help: ## Display this help screen
	@echo "Usage:"
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'