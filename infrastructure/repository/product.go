package repository

import (
	"gorm.io/gorm"

	"github.com/epessine/go-products-server/domain/model"
)

type Product struct {
	DB *gorm.DB
}

func (r Product) CreateProduct(product *model.Product) (*model.Product, error) {
	err := r.DB.Create(product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r Product) FindProducts() ([]model.Product, error) {
	var products []model.Product

	err := r.DB.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}
