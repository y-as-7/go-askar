package Services

import (
	"github.com/y-as-7/go-askar/app/Models"
	"github.com/y-as-7/go-askar/config"
)

type CustomerService struct{}

// NewService creates a new Customer service instance
func NewCustomerService() *CustomerService {
	return &CustomerService{}
}

// Create creates a new Customer
func (s *CustomerService) Create(data *Models.Customer) error {
	return config.DB.Create(data).Error
}

// GetByID retrieves a Customer by ID
func (s *CustomerService) GetByID(id uint) (*Models.Customer, error) {
	var item Models.Customer
	err := config.DB.First(&item, id).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

// GetAll retrieves all Customer records
func (s *CustomerService) GetAll() ([]Models.Customer, error) {
	var items []Models.Customer
	err := config.DB.Find(&items).Error
	return items, err
}

// Update updates a Customer
func (s *CustomerService) Update(id uint, data *Models.Customer) error {
	return config.DB.Model(&Models.Customer{}).Where("id = ?", id).Updates(data).Error
}

// Delete deletes a Customer by ID
func (s *CustomerService) Delete(id uint) error {
	return config.DB.Delete(&Models.Customer{}, id).Error
}
