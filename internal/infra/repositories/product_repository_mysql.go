package repository

import (
	"database/sql"

	entity "github.com/inaciogu/go-product-service/internal/entities"
)

type ProductRepositoryMysql struct {
	DB *sql.DB
}

func NewRepositoryMysql(db *sql.DB) *ProductRepositoryMysql {
	return &ProductRepositoryMysql{DB: db}
}

func (r *ProductRepositoryMysql) Create(product *entity.Product) error {
	_, err := r.DB.Exec("INSERT INTO products (id, name, price) VALUES (?, ?, ?)", product.ID, product.Name, product.Price)

	if err != nil {
		return err
	}
	return nil
}

func (r *ProductRepositoryMysql) FindAll() ([]*entity.Product, error) {
	rows, err := r.DB.Query("SELECT id, name, price FROM products")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*entity.Product

	for rows.Next() {
		var product entity.Product
		err = rows.Scan(&product.ID, &product.Name, &product.Price)

		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}
