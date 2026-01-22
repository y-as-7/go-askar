package Models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	
	Name     string `gorm:"not null" json:"name" binding:"required"`
	Email    string `gorm:"uniqueIndex;not null" json:"email" binding:"required,email"`
	Password string `gorm:"not null" json:"-"`
	Role     string `gorm:"type:varchar(50);default:'user'" json:"role"` // admin, user
	IsActive bool   `gorm:"default:true" json:"is_active"`
}

// HashPassword hashes the user password
func (u *User) HashPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// CheckPassword compares password with hashed password
func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

// BeforeCreate hook
func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.Role == "" {
		u.Role = "user"
	}
	return nil
}

// TableName specifies table name
func (User) TableName() string {
	return "users"
}
