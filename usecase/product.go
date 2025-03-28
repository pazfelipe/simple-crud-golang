package usecase

import "go-api/model"

type ProductUseCase struct {
}

func NewProductUseCase() ProductUseCase {
	return ProductUseCase{}
}

func (u *ProductUseCase) GetProducts() ([]model.Product, error) {
	return []model.Product{}, nil
}
