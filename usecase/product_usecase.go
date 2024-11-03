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

func (pu *ProductUsecase) CreateProduct(product entity.Product) (entity.Product, error) {
	productId, err := pu.repository.CreateProduct(product)

	if(err != nil) {
		return entity.Product{}, err
	}

	product.ID = productId

	return product, nil
}
