BIN_DIR=bin
APP_NAME=app
MAIN_PACKAGE=cmd/app

.DEFAULT_GOAL := help

help:
	@echo "Available commands:"
	@echo "  make build       - Build application"
	@echo "  make run         - Build and run"
	@echo "  make test        - Run all tests"
	@echo "  make test-coverage - Run tests with coverage"
	@echo "  make clean       - Clean build artifacts"
	@echo "  make deps        - Download dependencies"
	@echo "  make fmt         - Format code"
	@echo "  make vet         - Vet code"
	@echo "  make lint        - Lint code"
	@echo "  make check       - Run fmt, vet and test"
	@echo "  make build-prod  - Build production binary"

build:
	go build -o $(BIN_DIR)/$(APP_NAME) ./$(MAIN_PACKAGE)

run: build
	$(BIN_DIR)/$(APP_NAME)

test:
	go test ./...

test-coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

clean:
	if exist $(BIN_DIR) rmdir /s /q $(BIN_DIR)
	if exist coverage.out del coverage.out
	if exist coverage.html del coverage.html

deps:
	go mod download

fmt:
	go fmt ./...

vet:
	go vet ./...

lint:
	golangci-lint run

check: fmt vet test

build-prod:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o $(BIN_DIR)/$(APP_NAME) ./$(MAIN_PACKAGE)

.PHONY: help build run test test-coverage clean deps fmt vet lint check build-prod