package Services

import (
	"github.com/y-as-7/go-askar/app/Models"
	"github.com/y-as-7/go-askar/config"
)

type ProductService struct{}

// NewService creates a new Product service instance
func NewProductService() *ProductService {
	return &ProductService{}
}

// Create creates a new Product
func (s *ProductService) Create(data *Models.Product) error {
	return config.DB.Create(data).Error
}

// GetByID retrieves a Product by ID
func (s *ProductService) GetByID(id uint) (*Models.Product, error) {
	var item Models.Product
	err := config.DB.First(&item, id).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

// GetAll retrieves all Product records
func (s *ProductService) GetAll() ([]Models.Product, error) {
	var items []Models.Product
	err := config.DB.Find(&items).Error
	return items, err
}

// Update updates a Product
func (s *ProductService) Update(id uint, data *Models.Product) error {
	return config.DB.Model(&Models.Product{}).Where("id = ?", id).Updates(data).Error
}

// Delete deletes a Product by ID
func (s *ProductService) Delete(id uint) error {
	return config.DB.Delete(&Models.Product{}, id).Error
}
