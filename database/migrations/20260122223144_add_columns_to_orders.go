package migrations

import (
	"gorm.io/gorm"
)

// AddColumnsToOrders migration
type AddColumnsToOrders struct{}

// Up runs the migration
func (m *AddColumnsToOrders) Up(db *gorm.DB) error {
	// Add your migration logic here
	// Example:
	// return db.Exec("CREATE TABLE example (id INT PRIMARY KEY)").Error
	return nil
}

// Down reverses the migration
func (m *AddColumnsToOrders) Down(db *gorm.DB) error {
	// Add your rollback logic here
	// Example:
	// return db.Exec("DROP TABLE example").Error
	return nil
}
