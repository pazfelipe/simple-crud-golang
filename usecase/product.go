package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type ProductUseCase struct {
	repository repository.ProductRepository
}

func NewProductUseCase(repo repository.ProductRepository) ProductUseCase {
	return ProductUseCase{
		repository: repo,
	}
}

func (u *ProductUseCase) GetProducts() ([]model.Product, error) {
	return u.repository.GetProducts()
}

func (pu *ProductUseCase) CreateProduct(product model.Product) (model.Product, error) {
	productId, err := pu.repository.CreateProduct(product)

	if err != nil {
		return model.Product{}, err
	}

	product.ID = productId
	return product, nil
}

func (pu *ProductUseCase) GetProductById(productID int) (*model.Product, error) {
	product, err := pu.repository.GetProductById(productID)

	if err != nil {
		return nil, err
	}
	return product, nil
}

func (pu *ProductUseCase) DeleteProductById(productID int) error {
	err := pu.repository.DeleteProductById(productID)

	if err != nil {
		return err
	}
	return nil
}
