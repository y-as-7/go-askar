#!/bin/bash

# askar Framework Tool
# This script handles both CLI installation and project setup

set -e

# ANSI Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Check if we are inside an existing askar project (for setup mode)
IS_PROJECT=false
if [ -f "go.mod" ] && grep -q "github.com/y-as-7/go-askar" go.mod; then
    IS_PROJECT=true
fi

if [ "$IS_PROJECT" = false ]; then
    # --- CLI INSTALLER MODE ---
    echo -e "🚀 ${BLUE}askar CLI Installer${NC}"
    echo "==============================="
    
    # 1. Install Binary
    echo -e "📥 ${GREEN}Installing go-askar CLI...${NC}"
    go install github.com/y-as-7/go-askar@latest || { echo -e "${RED}❌ Failed to install go-askar. Make sure Go is installed.${NC}"; exit 1; }
    
    # 2. Configure Shell Alias (The Magic Trick)
    MAGIC_FUNC='go() { if [ "$1" == "askar" ]; then shift; go-askar "$@"; else command go "$@"; fi; }'
    
    configure_shell() {
        local shell_rc=$1
        if [ -f "$shell_rc" ]; then
            if ! grep -q "askar" "$shell_rc"; then
                echo -e "📝 ${GREEN}Adding 'go askar' magic trick to $shell_rc...${NC}"
                echo "" >> "$shell_rc"
                echo "# askar CLI shortcut" >> "$shell_rc"
                echo "$MAGIC_FUNC" >> "$shell_rc"
                return 0
            fi
        fi
        return 1
    }

    UPDATED=false
    configure_shell "$HOME/.bashrc" && UPDATED=true
    configure_shell "$HOME/.zshrc" && UPDATED=true

    echo ""
    echo -e "✅ ${GREEN}CLI installed successfully!${NC}"
    echo ""
    if [ "$UPDATED" = true ]; then
        echo -e "${YELLOW}IMPORTANT:${NC} Please restart your terminal or run:"
        echo -e "   ${BLUE}source ~/.bashrc${NC}  (or ~/.zshrc)"
    fi
    echo ""
    echo -e "Then you can start your project with just:"
    echo -e "   ${BLUE}go askar create/project my-app${NC}"
    echo ""
    exit 0
fi

# --- PROJECT SETUP MODE (Post-Clone) ---
echo -e "🚀 ${BLUE}askar Project Setup${NC}"
echo "==============================="
echo ""

# Get project name from arg or prompt
PROJECT_NAME=$1
if [ -z "$PROJECT_NAME" ]; then
    read -p "Enter your project name (e.g., my-shop): " PROJECT_NAME
fi

if [ -z "$PROJECT_NAME" ]; then
    echo -e "${RED}❌ Project name cannot be empty${NC}"
    exit 1
fi

echo -e "\n📦 ${GREEN}Setting up project: $PROJECT_NAME${NC}\n"

# Update go.mod
echo "📝 Updating go.mod..."
# Use a more robust sed for cross-platform (and avoid order-system specific match)
sed -i "s|module github.com/y-as-7/go-askar|module $PROJECT_NAME|g" go.mod || sed -i "s|module .*|module $PROJECT_NAME|g" go.mod

# Update all import paths
echo "📝 Updating import paths..."
find . -type f -name "*.go" -exec sed -i "s|github.com/y-as-7/go-askar/|$PROJECT_NAME/|g" {} +

# Create .env from .env.example
if [ ! -f .env ]; then
    echo "📝 Creating .env file..."
    cp .env.example .env
    JWT_SECRET=$(openssl rand -hex 32 2>/dev/null || cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w 64 | head -n 1)
    sed -i "s/CHANGE_THIS_TO_RANDOM_SECRET/$JWT_SECRET/" .env
    echo "✅ Generated random JWT_SECRET"
fi

# Install dependencies
echo "📥 Installing dependencies..."
go mod tidy

# Initialize git
if [ ! -d .git ] || [ "$(git rev-parse --is-inside-work-tree 2>/dev/null)" != "true" ]; then
    echo "🔧 Initializing git repository..."
    # If we are in a subfolder of a git repo (cloned for creation), we want a fresh start
    rm -rf .git || true
    git init
    git add .
    git commit -m "Initial commit: $PROJECT_NAME based on askar framework"
fi

echo -e "\n✅ ${GREEN}Project setup complete!${NC}\n"
echo -e "📋 ${BLUE}Next steps:${NC}"
echo -e "  1. Run: ${YELLOW}make run${NC}"
echo -e "  2. API at: ${YELLOW}http://localhost:8080${NC}\n"
echo -e "🎉 ${GREEN}Happy coding with askar!${NC}"
