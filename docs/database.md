# Database Guide

**askar** uses the powerful [GORM](https://gorm.io/) ORM for database management, providing a clean and expressive way to interact with your data.

## Supported Databases

Out of the box, askar supports:
- **SQLite** (Default, perfect for development)
- **PostgreSQL**
- **MySQL**

## Configuration

Configure your database connection in the `.env` file:

```env
DB_CONNECTION=sqlite
DB_HOST=127.0.0.1
DB_PORT=5432
DB_DATABASE=storage/database.db
DB_USERNAME=root
DB_PASSWORD=
```

For SQLite, only the `DB_DATABASE` path is required.

## Defining Models

Models are defined in the `app/models` directory. Here is an example of a simple `User` model:

```go
type User struct {
    gorm.Model
    Name     string `json:"name"`
    Email    string `gorm:"unique" json:"email"`
    Password string `json:"-"`
}
```

## Automatic Migrations

askar handles database migrations automatically. When you start the application (`make run`), GORM will automatically migrate your schema to match your models.

You can find the migration logic in `config/database.go`.

## Querying

Interact with your database using standard GORM syntax in your controllers:

```go
var user models.User
database.DB.First(&user, "email = ?", "test@example.com")
```

---

**Next Steps:**
- [Installation Guide](installation.md)
- [Authentication Guide](authentication.md)
