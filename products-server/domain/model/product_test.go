package model_test

import (
	"os"
	"testing"

	"github.com/epessine/go-products-server/domain/model"
)

func init() {
	os.Setenv("ENV", "test")
}

func TestModel_NewProduct(t *testing.T) {
	name := "Product Name"
	description := "Product Description"
	price := float32(123.12)

	product, err := model.NewProduct(name, description, price)
	if err != nil {
		t.Errorf("Expected no errors")
	}
	if product.Name != name {
		t.Errorf("Expected %s and got %s", name, product.Name)
	}
	if product.Description != description {
		t.Errorf("Expected %s and got %s", description, product.Description)
	}
	if product.Price != price {
		t.Errorf("Expected %e and got %e", price, product.Price)
	}
}