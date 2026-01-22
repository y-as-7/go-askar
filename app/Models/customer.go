package Models

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	
	// Add your fields here
}

// TableName specifies table name
func (Customer) TableName() string {
	return "customers"
}
