package repository

import (
	"database/sql"
	"fmt"
	"go-api/model"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "SELECT * FROM product"
	rows, err := pr.connection.Query(query)

	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}

	defer rows.Close()
	var products []model.Product
	var product model.Product

	for rows.Next() {
		err := rows.Scan(&product.ID, &product.Name, &product.Price)

		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}

		products = append(products, product)
	}
	return products, nil
}
