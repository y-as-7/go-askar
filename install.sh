#!/bin/bash

# go-askar Framework Installer
# This script sets up a new project from the go-askar framework template

set -e

echo "🚀 go-askar Framework Installer"
echo "==============================="
echo ""

# Get project name
read -p "Enter your project name (e.g., my-shop): " PROJECT_NAME

if [ -z "$PROJECT_NAME" ]; then
    echo "❌ Project name cannot be empty"
    exit 1
fi

# Validate project name
if ! [[ "$PROJECT_NAME" =~ ^[a-zA-Z0-9_-]+$ ]]; then
    echo "❌ Project name can only contain letters, numbers, hyphens, and underscores"
    exit 1
fi

echo ""
echo "📦 Setting up project: $PROJECT_NAME"
echo ""

# Update go.mod
echo "📝 Updating go.mod..."
sed -i "s/module order-system/module $PROJECT_NAME/" go.mod

# Update all import paths
echo "📝 Updating import paths..."
find . -type f -name "*.go" -exec sed -i "s|order-system/|$PROJECT_NAME/|g" {} +

# Create .env from .env.example
if [ ! -f .env ]; then
    echo "📝 Creating .env file..."
    cp .env.example .env
    
    # Generate random JWT secret
    JWT_SECRET=$(openssl rand -hex 32 2>/dev/null || cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w 64 | head -n 1)
    sed -i "s/CHANGE_THIS_TO_RANDOM_SECRET/$JWT_SECRET/" .env
    
    echo "✅ Generated random JWT_SECRET"
fi

# Install dependencies
echo "📥 Installing dependencies..."
go mod tidy

# Initialize git if not already initialized
if [ ! -d .git ]; then
    echo "🔧 Initializing git repository..."
    git init
    git add .
    git commit -m "Initial commit: $PROJECT_NAME based on go-askar framework"
fi

echo ""
echo "✅ Project setup complete!"
echo ""
echo "📋 Next steps:"
echo "  1. Review .env file and update configuration"
echo "  2. Run: make run"
echo "  3. API will be available at http://localhost:8080"
echo ""
echo "📚 Documentation:"
echo "  - Installation: docs/installation.md"
echo "  - Authentication: docs/authentication.md"
echo "  - Database: docs/database.md"
echo ""
echo "🎉 Happy coding with go-askar!"
