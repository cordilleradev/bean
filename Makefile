# Define variables for backend
BACKEND_DIR := backend
BINARY_NAME := app
GO_FILES := $(wildcard $(BACKEND_DIR)/*.go)

# Define variables for frontend
FRONTEND_DIR := frontend
FRONTEND_BUILD_DIR := $(FRONTEND_DIR)/dist

# Default target to build both frontend and backend
all: build-frontend build-backend

# Build the backend application
build-backend: $(GO_FILES)
	@echo "Building the Go application..."
	@cd $(BACKEND_DIR) && go build -o ../$(BINARY_NAME) main.go
	@echo "Backend build complete."

# Build the frontend application
build-frontend:
	@echo "Building the Vue.js application for production..."
	@cd $(FRONTEND_DIR) && npm install && npm run build
	@echo "Frontend build complete."

# Run the backend application
run-backend: build-backend
	@echo "Running the backend application..."
	@./$(BINARY_NAME) $(ARGS)

# Serve the frontend application
run-frontend: build-frontend
	@echo "Serving the frontend application..."
	@cd $(FRONTEND_BUILD_DIR)

# Serve the frontend application locally for development
run-frontend-local:
	@echo "Running the frontend application locally..."
	@cd $(FRONTEND_DIR) && npm install && npm run serve

# Clean up build artifacts
clean:
	@echo "Cleaning up..."
	@rm -f $(BINARY_NAME)
	@rm -rf $(FRONTEND_BUILD_DIR)
	@echo "Cleanup complete."

# Run tests for the backend
test-backend:
	@echo "Running backend tests..."
	@cd $(BACKEND_DIR) && go test ./...

# Format the backend code
fmt-backend:
	@echo "Formatting backend code..."
	@cd $(BACKEND_DIR) && go fmt ./...

# Lint the backend code
lint-backend:
	@echo "Linting backend code..."
	@cd $(BACKEND_DIR) && go vet ./...

.PHONY: all build-backend build-frontend run-backend serve-frontend clean test-backend fmt-backend lint-backend
