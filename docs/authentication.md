# Authentication

**askar** comes with a pre-configured, secure, JWT-based authentication system right out of the box.

## Overview

The authentication system uses:
- **JWT (JSON Web Tokens)**: For stateless session management.
- **bcrypt**: For secure password hashing.
- **Middleware**: To protect your API routes.

## Configuration

Your authentication settings are controlled via the `.env` file:

```env
JWT_SECRET=your_random_secret_here
JWT_EXPIRATION=24h
```

When you create a project using `go-askar create/project`, a random 64-character `JWT_SECRET` is automatically generated for you.

## Endpoints

The framework provides the following built-in authentication endpoints:

### 1. Registration
`POST /api/v1/auth/register`
Accepts `email`, `password`, and `name`.

### 2. Login
`POST /api/v1/auth/login`
Accepts `email` and `password`. Returns a JWT token.

### 3. Profile (Protected)
`GET /api/v1/auth/profile`
Requires an `Authorization: Bearer <token>` header.

## Protecting Routes

To protect your own routes, simply apply the JWT middleware in `routes/routes.go`:

```go
authGroup := v1.Group("/auth")
{
    // Public routes
    authGroup.POST("/register", controllers.Register)
    authGroup.POST("/login", controllers.Login)

    // Protected routes
    protected := authGroup.Group("/")
    protected.Use(middleware.JWTAuthMiddleware())
    {
        protected.GET("/profile", controllers.GetProfile)
    }
}
```

---

**Next Steps:**
- [Installation Guide](installation.md)
- [Database Guide](database.md)
