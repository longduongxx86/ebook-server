# Makefile for EBook Server

.PHONY: help build run test clean docker-build docker-up docker-down

# Default target
help:
	@echo "Available commands:"
	@echo "  build        - Build the application"
	@echo "  run          - Run the application locally"
	@echo "  test         - Run tests"
	@echo "  clean        - Clean build artifacts"
	@echo "  docker-build - Build Docker image"
	@echo "  docker-up    - Start services with Docker Compose"
	@echo "  docker-down  - Stop services with Docker Compose"
	@echo "  dev          - Run in development mode with hot reload"

# Build the application
build:
	go build -o bin/ebook-server main.go

# Run the application locally
run:
	go run main.go

# Run tests
test:
	go test ./...

# Clean build artifacts
clean:
	rm -rf bin/
	go clean

# Build Docker image
docker-build:
	docker build -t ebook-server .

# Start services with Docker Compose
docker-up:
	docker-compose up -d

# Stop services with Docker Compose
docker-down:
	docker-compose down

# Run in development mode (requires air for hot reload)
dev:
	@if command -v air > /dev/null; then \
		air; \
	else \
		echo "Air not installed. Installing..." && \
		go install github.com/cosmtrek/air@latest && \
		air; \
	fi

# Install dependencies
deps:
	go mod download
	go mod tidy

# Generate documentation
docs:
	@echo "API Documentation:"
	@echo "=================="
	@echo "Authentication:"
	@echo "  POST /api/v1/register - Register new user"
	@echo "  POST /api/v1/login    - Login user"
	@echo ""
	@echo "User Management:"
	@echo "  GET /api/v1/profile   - Get user profile (auth required)"
	@echo "  PUT /api/v1/profile  - Update user profile (auth required)"
	@echo ""
	@echo "Book Management:"
	@echo "  GET    /api/v1/books      - Get all books"
	@echo "  GET    /api/v1/books/:id  - Get book by ID"
	@echo "  POST   /api/v1/books     - Create book (auth required)"
	@echo "  PUT    /api/v1/books/:id  - Update book (auth required)"
	@echo "  DELETE /api/v1/books/:id  - Delete book (auth required)"
	@echo "  GET    /api/v1/my-books   - Get my books (auth required)"
	@echo ""
	@echo "Order Management:"
	@echo "  POST /api/v1/orders           - Create order (auth required)"
	@echo "  GET  /api/v1/orders          - Get my orders (auth required)"
	@echo "  GET  /api/v1/orders/:id      - Get order by ID (auth required)"
	@echo "  PUT  /api/v1/orders/:id/status - Update order status (auth required)"
