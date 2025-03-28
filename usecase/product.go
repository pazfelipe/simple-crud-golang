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
