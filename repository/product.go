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

func (pr *ProductRepository) CreateProduct(product model.Product) (int, error) {
	query, err := pr.connection.Prepare("INSERT INTO product(name, price) VALUES($1, $2) RETURNING id")
	var productID int

	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	defer query.Close()

	err = query.QueryRow(product.Name, product.Price).Scan(&productID)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return productID, nil
}

func (pr *ProductRepository) GetProductById(id int) (*model.Product, error) {
	query, err := pr.connection.Prepare("SELECT * FROM products WHERE id = $1;")
	var product model.Product

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer query.Close()
	err = query.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product with id %d not found", id)
		}

		return nil, err
	}
	return &product, nil
}
