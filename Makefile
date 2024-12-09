# Project Name
PROJECT_NAME := solcomp

# Directories
SRC_DIR := .
BUILD_DIR := ./bin
OUTPUT := $(BUILD_DIR)/$(PROJECT_NAME)

# Go Commands
GO := go
GOFMT := gofmt
GOTEST := $(GO) test
GOBUILD := $(GO) build
GOCLEAN := $(GO) clean
GOVET := $(GO) vet
GOMOD := $(GO) mod tidy

# Default Target
all: build

# Format Go Code
fmt:
	@echo "Formatting code..."
	@$(GOFMT) -w .

# Run Go Vet
vet:
	@echo "Running go vet..."
	@$(GOVET) ./...

# Run Tests
test:
	@echo "Running tests..."
	@$(GOTEST) ./...

# Tidy Go Modules
tidy:
	@echo "Tidying go modules..."
	@$(GOMOD)

# Build Project
build: tidy fmt vet
	@echo "Building $(PROJECT_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@$(GOBUILD) -o $(OUTPUT) $(SRC_DIR)

# Clean Build Artifacts
clean:
	@echo "Cleaning build artifacts..."
	@$(GOCLEAN)
	@rm -rf $(BUILD_DIR)

# Run Project
run: build
	@echo "Running $(PROJECT_NAME)..."
	@$(OUTPUT)

# Help
help:
	@echo "Available targets:"
	@echo "  all      - Default target: build the project"
	@echo "  fmt      - Format the Go code"
	@echo "  vet      - Run go vet on the code"
	@echo "  test     - Run tests"
	@echo "  tidy     - Tidy go.mod and go.sum"
	@echo "  build    - Build the project"
	@echo "  clean    - Clean build artifacts"
	@echo "  run      - Build and run the project"
	@echo "  help     - Display this help message"

.PHONY: all fmt vet test tidy build clean run help
