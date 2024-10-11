# Define variables
BACKEND_DIR := backend
BINARY_NAME := myapp
GO_FILES := $(wildcard $(BACKEND_DIR)/*.go)

# Default target to build the application
all: build

# Build the application
build: $(GO_FILES)
	@echo "Building the Go application..."
	@cd $(BACKEND_DIR) && go build -o ../$(BINARY_NAME) main.go
	@echo "Build complete."

# Run the application with arguments
run-backend: build
	@echo "Running the application with arguments: $(ARGS)"
	@./$(BINARY_NAME) $(ARGS)

# Clean up build artifacts
clean:
	@echo "Cleaning up..."
	@rm -f $(BINARY_NAME)
	@echo "Cleanup complete."

# Run tests
test:
	@echo "Running tests..."
	@cd $(BACKEND_DIR) && go test ./...

# Format the code
fmt:
	@echo "Formatting code..."
	@cd $(BACKEND_DIR) && go fmt ./...

# Lint the code
lint:
	@echo "Linting code..."
	@cd $(BACKEND_DIR) && go vet ./...

.PHONY: all build run-backend clean test fmt lint
