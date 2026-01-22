.PHONY: help run build clean test install migrate

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

install: ## Install dependencies
	go mod download
	go mod tidy

create: ## Create a new project (e.g., make create project=my-shop)
	@if [ -z "$(project)" ]; then echo "Error: project variable is required. usage: make create project=name"; exit 1; fi
	go run main.go create/project $(project)

install-cli: ## Install the askar CLI globally
	go install .

run: ## Run the application
	go run main.go

build: ## Build the application
	go build -o bin/go-askar main.go

clean: ## Clean build artifacts
	rm -rf bin/
	rm -f *.db

test: ## Run tests
	go test -v ./...

dev: ## Run with hot reload (requires air)
	air

migrate: ## Run database migrations
	@echo "Migrations run automatically on startup"

docker-build: ## Build Docker image
	docker build -t order-system:latest .

docker-run: ## Run Docker container
	docker run -p 8080:8080 order-system:latest

lint: ## Run linter
	golangci-lint run

format: ## Format code
	go fmt ./...
