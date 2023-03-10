package entity

import "github.com/google/uuid"

type ProductRepository interface {
	FindAll() ([]*Product, error)
	Create(product *Product) error
}

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	product := Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
	return &product
}
