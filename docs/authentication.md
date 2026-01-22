# Authentication Guide

go-askar comes with a complete authentication system using JWT tokens. This guide explains how to use and extend the authentication features.

## Overview

The authentication system provides:
- User registration with email and password
- Secure password hashing with bcrypt
- JWT token generation and validation
- Protected routes with middleware
- User profile management

## API Endpoints

### 1. Register a New User

**Endpoint**: `POST /api/v1/auth/register`

**Request Body**:
```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "password123"
}
```

**Response** (201 Created):
```json
{
  "success": true,
  "message": "User registered successfully",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "id": 1,
      "name": "John Doe",
      "email": "john@example.com",
      "role": "user",
      "is_active": true,
      "created_at": "2024-01-22T10:00:00Z",
      "updated_at": "2024-01-22T10:00:00Z"
    }
  }
}
```

**Example**:
```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "password123"
  }'
```

### 2. Login

**Endpoint**: `POST /api/v1/auth/login`

**Request Body**:
```json
{
  "email": "john@example.com",
  "password": "password123"
}
```

**Response** (200 OK):
```json
{
  "success": true,
  "message": "Login successful",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "id": 1,
      "name": "John Doe",
      "email": "john@example.com",
      "role": "user",
      "is_active": true
    }
  }
}
```

**Example**:
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "password123"
  }'
```

### 3. Get Profile (Protected)

**Endpoint**: `GET /api/v1/auth/profile`

**Headers**: 
```
Authorization: Bearer YOUR_JWT_TOKEN
```

**Response** (200 OK):
```json
{
  "success": true,
  "data": {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com",
    "role": "user",
    "is_active": true,
    "created_at": "2024-01-22T10:00:00Z",
    "updated_at": "2024-01-22T10:00:00Z"
  }
}
```

**Example**:
```bash
TOKEN="your_jwt_token_here"
curl http://localhost:8080/api/v1/auth/profile \
  -H "Authorization: Bearer $TOKEN"
```

## How It Works

### Password Security

Passwords are hashed using bcrypt before being stored in the database:

```go
// In app/models/user.go
func (u *User) HashPassword(password string) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    u.Password = string(hashedPassword)
    return nil
}
```

### JWT Tokens

JWT tokens are generated upon successful registration or login:

- **Payload**: Contains user ID, email, and role
- **Expiry**: Configured in `.env` (default: 24 hours)
- **Secret**: Must be set in `.env` as `JWT_SECRET`

### Protected Routes

To protect a route, use the `AuthMiddleware()`:

```go
// In routes/routes.go
protected := v1.Group("")
protected.Use(middleware.AuthMiddleware())
{
    protected.GET("/auth/profile", authCtrl.GetProfile)
    // Add more protected routes here
}
```

### Admin-Only Routes

For admin-only endpoints, use both middlewares:

```go
admin := v1.Group("")
admin.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
{
    // Admin routes here
}
```

## Customization

### Change JWT Expiry

Edit `.env`:
```env
JWT_EXPIRY_HOURS=48  # Token valid for 48 hours
```

### Add Custom User Fields

1. Edit `app/models/user.go`:
```go
type User struct {
    // ... existing fields
    Phone     string `gorm:"type:varchar(20)" json:"phone"`
    CompanyID uint   `json:"company_id"`
}
```

2. The migration will run automatically on next start

### Add More Roles

1. Update the User model:
```go
Role string `gorm:"type:varchar(50);default:'user'" json:"role"` // admin, user, manager
```

2. Create custom middleware:
```go
func ManagerMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        role, _ := c.Get("user_role")
        if role != "manager" && role != "admin" {
            c.JSON(403, dto.Response{
                Success: false,
                Error:   "Manager access required",
            })
            c.Abort()
            return
        }
        c.Next()
    }
}
```

## Security Best Practices

1. **Always use HTTPS in production**
2. **Set a strong JWT_SECRET** (64+ characters)
3. **Set GIN_MODE=release in production**
4. **Implement rate limiting** for auth endpoints
5. **Add email verification** for production use
6. **Implement password reset** functionality
7. **Rotate JWT secrets** periodically
8. **Log authentication attempts**

## Error Codes

| Status | Error | Description |
|--------|-------|-------------|
| 400 | Bad Request | Invalid request data |
| 401 | Unauthorized | Invalid credentials or missing token |
| 403 | Forbidden | Insufficient permissions |
| 409 | Conflict | Email already exists |
| 500 | Internal Server Error | Server error |

## Testing with Postman

1. Import the API collection (if available)
2. Register a user
3. Use the returned token for protected requests
4. Set token in Authorization → Bearer Token

## Next Steps

- [Database Guide](database.md) - Add more models
- Extend authentication with email verification
- Add password reset functionality
- Implement OAuth providers
