# Go parameters
APP_NAME := Link_Scraper
MAIN_FILE := main.go
BUILD_DIR := bin
GO := go

# Default target
.PHONY: all
all: build

# Build the application
.PHONY: build
build:
	@echo ">> Building $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@$(GO) build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_FILE)

# Run the application
.PHONY: run
run:
	@echo ">> Running $(APP_NAME)..."
	@$(GO) run $(MAIN_FILE)

# Run tests with coverage
.PHONY: test
test:
	@echo ">> Running tests..."
	@$(GO) test ./... -v -cover

# Format code
.PHONY: fmt
fmt:
	@echo ">> Formatting code..."
	@$(GO) fmt ./...

# Tidy up dependencies
.PHONY: tidy
tidy:
	@echo ">> Tidying modules..."
	@$(GO) mod tidy

# Clean build artifacts
.PHONY: clean
clean:
	@echo ">> Cleaning build artifacts..."
	@rm -rf $(BUILD_DIR)

# Cross-compile for Linux (example)
.PHONY: build-linux
build-linux:
	@echo ">> Cross-compiling for Linux..."
	@mkdir -p $(BUILD_DIR)
	@GOOS=linux GOARCH=amd64 $(GO) build -o $(BUILD_DIR)/$(APP_NAME)-linux $(MAIN_FILE)