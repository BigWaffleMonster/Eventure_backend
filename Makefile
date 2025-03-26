# Variables
GO := go
APP_NAME := ""
BINARY := bin/$(APP_NAME)
TEST_REPORT := test-report.xml

# Default target
.PHONY: all
all: build

# Build the application
.PHONY: build
build:
	@echo "Building $(APP_NAME)..."
	mkdir -p bin
	$(GO) build -o $(BINARY) ./cmd/$(APP_NAME)
	@echo "Build complete: $(BINARY)"

# Run the application
.PHONY: run
run:
	@echo "Running $(APP_NAME)..."
	$(GO) run ./cmd/$(APP_NAME)

# Install dependencies
.PHONY: deps
deps:
	@echo "Installing dependencies..."
	$(GO) mod tidy
	@echo "Dependencies installed."

# Format the code
.PHONY: fmt
fmt:
	@echo "Formatting code..."
	$(GO) fmt ./...
	@echo "Code formatting complete."

# Lint the code
.PHONY: lint
lint:
	@echo "Linting code..."
	golangci-lint run ./...
	@echo "Linting complete."

# Run tests
.PHONY: test
test:
	@echo "Running tests..."
	$(GO) test -v ./... -coverprofile=coverage.out
	@echo "Tests complete."

# Generate test report in JUnit format (optional)
.PHONY: test-report
test-report:
	@echo "Generating test report..."
	go install github.com/jstemmer/go-junit-report@latest
	$(GO) test -v ./... | go-junit-report > $(TEST_REPORT)
	@echo "Test report generated: $(TEST_REPORT)"

# Clean up build artifacts
.PHONY: clean
clean:
	@echo "Cleaning up..."
	rm -rf bin/
	rm -f coverage.out
	@echo "Cleanup complete."

# Run all checks (format, lint, test)
.PHONY: check
check: fmt lint test

# Help command to display available targets
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  all         - Build the application (default)"
	@echo "  build       - Build the application"
	@echo "  run         - Run the application"
	@echo "  deps        - Install dependencies"
	@echo "  fmt         - Format the code"
	@echo "  lint        - Lint the code"
	@echo "  test        - Run tests"
	@echo "  test-report - Generate test report in JUnit format"
	@echo "  clean       - Clean up build artifacts"
	@echo "  check       - Run all checks (format, lint, test)"
	@echo "  help        - Display this help message"
