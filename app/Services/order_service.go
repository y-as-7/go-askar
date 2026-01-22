package Services

import (
	"github.com/y-as-7/go-askar/app/Models"
	"github.com/y-as-7/go-askar/config"
)

type OrderService struct{}

// NewService creates a new Order service instance
func NewOrderService() *OrderService {
	return &OrderService{}
}

// Create creates a new Order
func (s *OrderService) Create(data *Models.Order) error {
	return config.DB.Create(data).Error
}

// GetByID retrieves a Order by ID
func (s *OrderService) GetByID(id uint) (*Models.Order, error) {
	var item Models.Order
	err := config.DB.First(&item, id).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

// GetAll retrieves all Order records
func (s *OrderService) GetAll() ([]Models.Order, error) {
	var items []Models.Order
	err := config.DB.Find(&items).Error
	return items, err
}

// Update updates a Order
func (s *OrderService) Update(id uint, data *Models.Order) error {
	return config.DB.Model(&Models.Order{}).Where("id = ?", id).Updates(data).Error
}

// Delete deletes a Order by ID
func (s *OrderService) Delete(id uint) error {
	return config.DB.Delete(&Models.Order{}, id).Error
}
