package model

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type ProductRepository interface {
	CreateProduct(product *Product) (*Product, error)
	FindProducts() ([]Product, error)
}

type Product struct {
	ID          uint
	Name        string  `validate:"required"`
	Description string  `validate:"required"`
	Price       float32 `validate:"required,gt=0"`
	CreatedAt   time.Time
}

var validate *validator.Validate

func init() {
	validate = validator.New(validator.WithRequiredStructEnabled())
}

func (product *Product) isValid() error {
	err := validate.Struct(product)
	if err != nil {
		return err
	}
	return nil
}

func NewProduct(name string, description string, price float32) (*Product, error) {
	product := Product{
		Name:        name,
		Description: description,
		Price:       price,
	}

	err := product.isValid()
	if err != nil {
		return nil, err
	}

	return &product, nil
}
