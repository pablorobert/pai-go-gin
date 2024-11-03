package repository

import (
	"database/sql"
	"fmt"
	"go-api/entity"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

func (pr *ProductRepository) GetProducts() ([]entity.Product, error) {
	query := "SELECT id_product, name, price FROM products"
	rows, err := pr.connection.Query(query)

	if err != nil {
		fmt.Println(err)

		return []entity.Product{}, err
	}

	var productList []entity.Product
	var productObjt entity.Product

	for rows.Next() {
		err = rows.Scan(
			&productObjt.ID,
			&productObjt.Name,
			&productObjt.Price,
		)
		if err != nil {
			fmt.Println(err)
			return []entity.Product{}, err
		}

		productList = append(productList, productObjt)
	}

	rows.Close()

	return productList, nil

}


func (pr *ProductRepository) CreateProduct(product entity.Product) (int, error)  {
	var id_product int
	query, err := pr.connection.Prepare("INSERT INTO products(name, price) VALUES ($1, $2) RETURNING id_product")

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(product.Name, product.Price).Scan(&id_product)
	if err != nil {
		fmt.Println(err)
		return 0, err
	} 

	query.Close()

	return id_product, nil


}
