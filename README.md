# askar

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/License-MIT-green.svg?style=for-the-badge)](LICENSE)
[![Framework](https://img.shields.io/badge/Framework-Gin-00ADD8?style=for-the-badge)](https://gin-gonic.com/)
[![GORM](https://img.shields.io/badge/ORM-GORM-00ADD8?style=for-the-badge)](https://gorm.io/)

## About askar

askar is a Laravel-inspired Go web application framework with expressive, elegant syntax. We believe development must be an enjoyable and creative experience to be truly fulfilling. askar takes the pain out of development by easing common tasks used in many web projects, such as:

- **Built-in Authentication System** - Complete JWT-based user authentication out of the box
- **Multi-Database Support** - SQLite, PostgreSQL, and MySQL support
- **Powerful ORM** - GORM integration with automatic migrations
- **Secure by Default** - bcrypt password hashing and JWT token management
- **Clean Architecture** - Well-organized MVC structure
- **One-Command Installation** - Get your API running in seconds
- **Comprehensive Documentation** - Detailed guides for every feature

askar is accessible, powerful, and provides tools required for building robust RESTful APIs.

## Installation

Install the CLI tool:

```bash
go install github.com/y-as-7/askar@latest
```

Create a new project:

```bash
go run main.go create/project my-project
```
or if installed globally:
```bash
askar create/project my-project
```

That's it! Just like Laravel's `composer create-project`, one simple command creates your entire project with:
- 🎨 Beautiful 3D logo display
- 📦 Automatic framework download
- ⚙️ Complete configuration
- 🔐 Secure JWT secret generation
- 📥 Dependency installation

## Quick Start

After creating your project:

```bash
cd my-project
cp .env.example .env
make run
```

Your API will be available at `http://localhost:8080`

## Learning askar

askar has comprehensive documentation to help you get started:

- **[Installation Guide](docs/installation.md)** - Complete setup instructions
- **[Authentication](docs/authentication.md)** - User authentication and JWT tokens
- **[Database Guide](docs/database.md)** - Working with models and migrations

## Project Structure

```
my-project/
├── app/
│   ├── controllers/       # Request handlers
│   ├── middleware/        # JWT authentication
│   ├── models/           # Database models
│   └── dto/              # Data transfer objects
├── config/               # Configuration files
├── routes/               # API route definitions
├── storage/              # Database and logs
├── docs/                 # Documentation
└── main.go              # Application entry
```

## API Endpoints

### Public Endpoints
```
POST /api/v1/auth/register   # User registration
POST /api/v1/auth/login      # User login
GET  /health                 # Health check
```

### Protected Endpoints
```
GET /api/v1/auth/profile     # Get user profile (requires JWT)
```

## CLI Tool Features

The `askar` command provides a beautiful installation experience:

```
   ██████╗  ██████╗        █████╗ ███████╗██╗  ██╗ █████╗ ██████╗ 
  ██╔════╝ ██╔═══██╗      ██╔══██╗██╔════╝██║ ██╔╝██╔══██╗██╔══██╗
  ██║  ███╗██║   ██║█████╗███████║███████╗█████╔╝ ███████║██████╔╝
  ██║   ██║██║   ██║╚════╝██╔══██║╚════██║██╔═██╗ ██╔══██║██╔══██╗
  ╚██████╔╝╚██████╔╝      ██║  ██║███████║██║  ██╗██║  ██║██║  ██║
   ╚═════╝  ╚═════╝       ╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═╝

              A Laravel-inspired Go Framework
```

- 🎨 Colorful terminal output
- ⚡ Animated loading indicators
- ✨ 3D ASCII art logo
- 📦 Automatic project setup

## Contributing

Thank you for considering contributing to askar! The contribution guide can be found in the [CONTRIBUTING.md](CONTRIBUTING.md) document.

## Security Vulnerabilities

If you discover a security vulnerability within askar, please send an email to the maintainers. All security vulnerabilities will be promptly addressed.

**Security Best Practices:**
- Always use HTTPS in production
- Set a strong `JWT_SECRET` (64+ characters)
- Set `GIN_MODE=release` in production
- Implement rate limiting for authentication endpoints
- Keep dependencies up to date

## Code of Conduct

askar is committed to providing a welcoming and inclusive environment for all contributors. Please be respectful and constructive in all interactions.

## License

The askar framework is open-sourced software licensed under the [MIT license](LICENSE).

---

**Built with ❤️ for the Go community**
