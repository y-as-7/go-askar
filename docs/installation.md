# Installation Guide

This guide will walk you through installing and setting up a new project using the go-askar framework.

## Prerequisites

- Go 1.21 or higher
- Git
- (Optional) PostgreSQL or MySQL if not using SQLite

## Quick Installation

### 1. Install the CLI Tool

```bash
go install github.com/y-as-7/go-askar@latest
```

Make sure `$GOPATH/bin` is in your PATH.

### 2. Create a New Project

```bash
askar new my-project
```

The CLI will display a beautiful interface and automatically:
- Clone the framework
- Configure your project
- Update all imports
- Generate JWT secret
- Install dependencies
- Initialize git repository

**Output:**
```
   ██████╗  ██████╗        █████╗ ███████╗██╗  ██╗ █████╗ ██████╗ 
  ██╔════╝ ██╔═══██╗      ██╔══██╗██╔════╝██║ ██╔╝██╔══██╗██╔══██╗
  ██║  ███╗██║   ██║█████╗███████║███████╗█████╔╝ ███████║██████╔╝
  ██║   ██║██║   ██║╚════╝██╔══██║╚════██║██╔═██╗ ██╔══██║██╔══██╗
  ╚██████╔╝╚██████╔╝      ██║  ██║███████║██║  ██╗██║  ██║██║  ██║
   ╚═════╝  ╚═════╝       ╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═╝

              A Laravel-inspired Go Framework
              Version 1.0.0

▸ 🚀 Creating new go-askar project: my-project
▸ 📦 Downloading framework...
  ✓ Repository cloned
▸ 📝 Configuring project...
▸ 🔧 Updating imports...
▸ ⚙️  Creating environment file...
▸ 📥 Installing dependencies...
▸ 🔧 Initializing git repository...

✓ Project created successfully!

╔════════════════════════════════════════════════════════════╗
║                                                            ║
║  ✓ Your go-askar project is ready!                        ║
║                                                            ║
║  Next steps:                                               ║
║                                                            ║
║    cd my-project                                           ║
║    cp .env.example .env                                    ║
║    make run                                                ║
║                                                            ║
║  Documentation: docs/                                      ║
║  API: http://localhost:8080                                ║
║                                                            ║
╚════════════════════════════════════════════════════════════╝

Happy coding! 🎉
```

### 3. Run Your Project

```bash
cd my-project
cp .env.example .env
make run
```

Your API will be running at `http://localhost:8080`!

## Alternative: Manual Installation

If you prefer not to use the CLI tool:

```bash
git clone https://github.com/y-as-7/go-askar my-project
cd my-project
./install.sh
```

## Configuration

### Environment Setup

Edit the `.env` file to customize your configuration:

```bash
nano .env  # or use your preferred editor
```

**Important**: The JWT_SECRET is generated automatically, but you can change it if needed.

### Database Selection

#### SQLite (Default - Recommended for Development)

No additional setup required! The database file will be created automatically at `storage/database/app.db`.

#### PostgreSQL

1. Create a database:
```bash
createdb your_database_name
```

2. Update `.env`:
```env
DB_DRIVER=postgres
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=your_database_name
DB_SSLMODE=disable
```

#### MySQL

1. Create a database:
```sql
CREATE DATABASE your_database_name;
```

2. Update `.env`:
```env
DB_DRIVER=mysql
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=your_database_name
```

## Verification

Test the installation:

```bash
curl http://localhost:8080/health
```

You should see:
```json
{
  "success": true,
  "message": "API is running"
}
```

## Available Commands

```bash
make install    # Install dependencies
make run        # Run the application
make build      # Build binary
make clean      # Clean build artifacts
make test       # Run tests
make format     # Format code
```

## Troubleshooting

### "askar: command not found"

Make sure `$GOPATH/bin` is in your PATH:

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

Add this to your `~/.bashrc` or `~/.zshrc` to make it permanent.

### Permission denied when running install.sh

Make the script executable:
```bash
chmod +x install.sh
```

### Database connection errors

- Verify your database is running
- Check credentials in `.env`
- Ensure the database exists

### Port already in use

Change the `PORT` in your `.env` file:
```env
PORT=3000
```

## Next Steps

- [Authentication Guide](authentication.md) - Learn how to use the auth system
- [Database Guide](database.md) - Working with models and migrations
- Start building your API!

## Support

If you encounter any issues, please check the [GitHub Issues](https://github.com/y-as-7/go-askar/issues).
