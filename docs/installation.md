# Installation Guide

Welcome to **askar**! This guide will walk you through setting up the framework and creating your first project.

## Prerequisites

- [Go](https://go.dev/doc/install) 1.21 or higher
- [Git](https://git-scm.com/downloads)
- A terminal of your choice

## 1. Install the CLI Tool

The **askar** CLI is the primary tool for creating new projects. Install it using the standard Go installer:

```bash
go install github.com/y-as-7/go-askar@latest
```

> [!NOTE]
> This will install a binary named `go-askar`.

### PATH Troubleshooting
If your terminal says `go-askar: command not found`, you need to add your Go binary directory to your system's PATH:

```bash
# Temporarily (current session)
export PATH=$PATH:$(go env GOPATH)/bin

# Permanently (add to ~/.bashrc or ~/.zshrc)
echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.bashrc
source ~/.bashrc
```

## 2. Create Your First Project

Once the CLI is installed, you can create a new project with a single command:

```bash
go-askar create/project my-new-app
```

This will:
- Clone the latest version of the askar framework.
- Configure your project naming and internal imports.
- Generate a secure `JWT_SECRET` for authentication.
- Initialize a fresh Git repository for your project.

## 3. ✨ The "go askar" Pro Tip

Many developers prefer to run the command as **`go askar`**. While Go doesn't natively support custom subcommands, you can enable this syntax easily by adding a shell function to your `~/.bashrc` (or `~/.zshrc`):

```bash
# Add to your shell profile
go() {
  if [ "$1" == "askar" ]; then
    shift
    go-askar "$@"
  else
    command go "$@"
  fi
}
```

After running `source ~/.bashrc`, you can simply run:
```bash
go askar create/project my-project
```

## 4. Getting Started

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
