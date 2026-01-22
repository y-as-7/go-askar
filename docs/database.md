# Database Guide

go-askar uses GORM as its ORM, supporting SQLite, PostgreSQL, and MySQL. This guide covers database configuration, models, and migrations.

## Database Configuration

### Supported Databases

- **SQLite** - Default, no setup required
- **PostgreSQL** - Production-ready relational database
- **MySQL** - Alternative production database

### Configuration

Edit your `.env` file:

#### SQLite (Default)
```env
DB_DRIVER=sqlite
DB_PATH=storage/database/app.db
```

#### PostgreSQL
```env
DB_DRIVER=postgres
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=your_database
DB_SSLMODE=disable
```

#### MySQL
```env
DB_DRIVER=mysql
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=your_database
```

## Migrations

go-askar uses GORM's AutoMigrate feature. Migrations run automatically when the application starts.

###Current Migrations

The framework includes:
- `users` table with authentication fields

### Adding Models to Migration

Edit `config/database.go`:

```go
func AutoMigrate() error {
    return DB.AutoMigrate(
        &models.User{},
        &models.Product{},     // Add your model here
        &models.Order{},       // Add another model
    )
}
```

## Creating Models

### Basic Model

Create a new file in `app/models/`:

```go
// app/models/product.go
package models

import (
    "time"
    "gorm.io/gorm"
)

type Product struct {
    ID        uint           `gorm:"primarykey" json:"id"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
    
    Name        string  `gorm:"not null" json:"name"`
    Description string  `gorm:"type:text" json:"description"`
    Price       float64 `gorm:"not null" json:"price"`
    Stock       int     `gorm:"default:0" json:"stock"`
    IsActive    bool    `gorm:"default:true" json:"is_active"`
}

func (Product) TableName() string {
    return "products"
}
```

### Model with Relationships

```go
// app/models/order.go
package models

import (
    "time"
    "gorm.io/gorm"
)

type Order struct {
    ID        uint           `gorm:"primarykey" json:"id"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
    
    UserID      uint        `gorm:"not null" json:"user_id"`
    User        User        `gorm:"foreignKey:UserID" json:"user"`
    TotalAmount float64     `gorm:"not null" json:"total_amount"`
    Status      string      `gorm:"type:varchar(50);default:'pending'" json:"status"`
    OrderItems  []OrderItem `gorm:"foreignKey:OrderID" json:"order_items,omitempty"`
}

type OrderItem struct {
    ID        uint    `gorm:"primarykey" json:"id"`
    OrderID   uint    `gorm:"not null" json:"order_id"`
    ProductID uint    `gorm:"not null" json:"product_id"`
    Product   Product `gorm:"foreignKey:ProductID" json:"product"`
    Quantity  int     `gorm:"not null" json:"quantity"`
    Price     float64 `gorm:"not null" json:"price"`
}
```

## GORM Field Tags

Common field tags:

```go
type Example struct {
    // Primary key
    ID uint `gorm:"primarykey"`
    
    // Not null constraint
    Name string `gorm:"not null"`
    
    // Unique constraint
    Email string `gorm:"uniqueIndex"`
    
    // Default value
    Status string `gorm:"default:'active'"`
    
    // Column type
    Bio string `gorm:"type:text"`
    
    // Foreign key
    UserID uint `gorm:"not null"`
    User   User `gorm:"foreignKey:UserID"`
    
    // JSON tag (for API responses)
    Price float64 `json:"price"`
    
    // Ignore field
    Password string `json:"-"`
}
```

## Database Operations

### Create

```go
// In your controller
product := models.Product{
    Name:        "iPhone 15",
    Description: "Latest model",
    Price:       999.99,
    Stock:       50,
}

if err := config.DB.Create(&product).Error; err != nil {
    // Handle error
}
```

### Read

```go
// Find by ID
var product models.Product
config.DB.First(&product, id)

// Find all
var products []models.Product
config.DB.Find(&products)

// With conditions
config.DB.Where("price > ?", 500).Find(&products)

// With preloading (relationships)
config.DB.Preload("OrderItems.Product").Find(&orders)
```

### Update

```go
// Update specific fields
config.DB.Model(&product).Updates(models.Product{
    Price: 899.99,
    Stock: 45,
})

// Update single field
config.DB.Model(&product).Update("stock", 40)
```

### Delete

```go
// Soft delete (sets DeletedAt)
config.DB.Delete(&product, id)

// Permanent delete
config.DB.Unscoped().Delete(&product, id)
```

## Accessing the Database

The database instance is available globally:

```go
import "your-project/config"

// In your function
config.DB.Where("email = ?", email).First(&user)
```

## Advanced Queries

### Joins

```go
var results []struct {
    UserName    string
    OrderCount  int64
}

config.DB.Table("users").
    Select("users.name as user_name, count(orders.id) as order_count").
    Joins("left join orders on orders.user_id = users.id").
    Group("users.id").
    Scan(&results)
```

### Transactions

```go
err := config.DB.Transaction(func(tx *gorm.DB) error {
    // Create order
    if err := tx.Create(&order).Error; err != nil {
        return err
    }
    
    // Create order items
    for _, item := range orderItems {
        if err := tx.Create(&item).Error; err != nil {
            return err
        }
    }
    
    return nil
})
```

## Best Practices

1. **Always handle errors** from database operations
2. **Use transactions** for related operations
3. **Index frequently queried fields**
4. **Use soft deletes** for important data
5. **Preload relationships** to avoid N+1 queries
6. **Validate data** before database operations
7. **Use prepared statements** (GORM does this automatically)

## Troubleshooting

### Migration doesn't run

Check `config/database.go` - ensure your model is in `AutoMigrate()`.

### Foreign key errors

Ensure parent record exists before creating child records.

### Connection pool exhaustion

Configure connection pool in `config/database.go`:

```go
sqlDB, _ := DB.DB()
sqlDB.SetMaxIdleConns(10)
sqlDB.SetMaxOpenConns(100)
sqlDB.SetConnMaxLifetime(time.Hour)
```

## References

- [GORM Documentation](https://gorm.io/docs/)
- [SQLite](https://www.sqlite.org/docs.html)
- [PostgreSQL](https://www.postgresql.org/docs/)
- [MySQL](https://dev.mysql.com/doc/)
