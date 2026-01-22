package migrations

import (
	"gorm.io/gorm"
)

// CreateProductsTable migration
type CreateProductsTable struct{}

// Up runs the migration
func (m *CreateProductsTable) Up(db *gorm.DB) error {
	// Add your migration logic here
	// Example:
	// return db.Exec("CREATE TABLE example (id INT PRIMARY KEY)").Error
	return nil
}

// Down reverses the migration
func (m *CreateProductsTable) Down(db *gorm.DB) error {
	// Add your rollback logic here
	// Example:
	// return db.Exec("DROP TABLE example").Error
	return nil
}
