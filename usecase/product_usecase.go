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

	if err != nil {
		return entity.Product{}, err
	}

	product.ID = productId

	return product, nil
}

func (pu *ProductUsecase) GetProductById(id_product int) (*entity.Product, error) {
	product, err := pu.repository.GetProductById(id_product)

	if err != nil {
		return nil, err
	}
	return product, nil
}

func (pu *ProductUsecase) UpdateProduct(product entity.Product) (*entity.Product, error) {
	productId, err := pu.repository.UpdateProduct(product)

	if err != nil {
		return nil, err
	}

	// Verifica se o produto foi encontrado e retornado

	/*if productId == nil {
		return nil, fmt.Errorf("produto com ID %d nao encontrado", product.ID)
	}*/

	return productId, nil
}

func (pu *ProductUsecase) DeleteProductById(id_product int) error {
	err := pu.repository.DeleteProductById(id_product)

	if err != nil {
		return err
	}
	return nil
}
