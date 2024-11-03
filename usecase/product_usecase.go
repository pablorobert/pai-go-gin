package usecase

import (
	"go-api/entity"
	"go-api/repository"
)

type ProductUsecase struct {
	repository repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repo,
	}
}

func (pu *ProductUsecase) GetProducts() ([]entity.Product, error) {
	return pu.repository.GetProducts()
}
