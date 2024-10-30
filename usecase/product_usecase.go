package usecase 


import "go-api/entity"

type ProductUsecase struct {
	// repository
}


func NewProductUsecase() ProductUsecase{
	return ProductUsecase{}
}


func(pu *ProductUsecase) GetProducts()([]entity.Product, error) {
	return []entity.Product{}, nil
}
