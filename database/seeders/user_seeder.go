package seeders

import (
	"github.com/y-as-7/go-askar/app/Models"
	"gorm.io/gorm"
)

type UserSeeder struct{}

func (s *UserSeeder) Run(db *gorm.DB) error {
	// Add your seeding logic here
	// Example:
	// return db.Create(&Models.User{Name: "Admin", Email: "admin@example.com"}).Error
	return nil
}
