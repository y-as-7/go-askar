# Installation Guide

Welcome to **askar**! This guide will walk you through setting up the framework and creating your first project using our simplified 2-command process.

## 1. One-Line Installation

Choose the command for your platform to install the **askar** CLI and configure your terminal automatically:

### 🍎 macOS / 🐧 Linux
```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/y-as-7/go-askar/main/install.sh)"
```

### 🪟 Windows (PowerShell)
```powershell
iwr -useb https://raw.githubusercontent.com/y-as-7/go-askar/main/install.ps1 | iex
```

### What this does:
- Installs the `go-askar` CLI binary.
- Automatically adds a "Magic Trick" (shell function or alias) to your terminal profile.
- Enables the **`go askar`** command syntax.

> [!IMPORTANT]
> After running the command, **restart your terminal** or run `source ~/.bashrc` (or `. $PROFILE` on Windows) to apply the changes.

## 2. Create Your First Project

Once installed, you can create a new project using the exact syntax below:

```bash
go askar create/project my-new-app
```

This command will:
- Download the framework.
- Configure all naming and internal imports.
- Generate a secure `JWT_SECRET`.
- Ready for your own `git init`.

## 3. Getting Started

After creating your project, jump into the directory and start the server:

```bash
cd my-new-app
cp .env.example .env
make run
```

Your API is now live at `http://localhost:8080`!

---

**Next Steps:**
- Learn about [Authentication](authentication.md)
- Configure your [Database](database.md)
