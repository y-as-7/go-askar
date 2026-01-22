package seeders

import (
	"gorm.io/gorm"
)

type ProductSeeder struct{}

func (s *ProductSeeder) Run(db *gorm.DB) error {
	// Add your seeding logic here
	// Example:
	// return db.Create(&Models.User{Name: "Admin", Email: "admin@example.com"}).Error
	return nil
}
