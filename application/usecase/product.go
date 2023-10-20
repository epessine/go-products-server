package usecase

import "github.com/epessine/go-products-server/domain/model"

type Product struct {
	Repository model.ProductRepository
}

func (p *Product) CreateProduct(name string, description string, price float32) (*model.Product, error) {
	product, err := model.NewProduct(name, description, price)
	if err != nil {
		return nil, err
	}

	product, err = p.Repository.CreateProduct(product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *Product) FindProducts() ([]model.Product, error) {
	products, err := p.Repository.FindProducts()
	if err != nil {
		return nil, err
	}
	return products, err
}
