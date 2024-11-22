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

		return nil, err
	}

	productList := []entity.Product{}
	var productObjt entity.Product

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&productObjt.ID,
			&productObjt.Name,
			&productObjt.Price,
		)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		productList = append(productList, productObjt)
	}

	return productList, nil

}

func (pr *ProductRepository) CreateProduct(product entity.Product) (int, error) {
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

func (pr *ProductRepository) GetProductById(id_product int) (*entity.Product, error) {
	query, err := pr.connection.Prepare("SELECT * FROM products WHERE id_product = $1")

	defer query.Close()

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var produto entity.Product

	err = query.QueryRow(id_product).Scan(
		&produto.ID,
		&produto.Name,
		&produto.Price,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}

		return nil, err
	}
	return &produto, nil
}

func (pr *ProductRepository) UpdateProduct(product entity.Product) (*entity.Product, error) {
	query, err := pr.connection.Prepare("UPDATE products SET name=$1, price=$2 WHERE id_product=$3")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// fechando a comunicao com o banco
	defer query.Close()

	// Executando a atualização
	_, err = query.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		return nil, err
	}
	// Se a atualização foi bem-sucedida, não retornamos um produto, apenas confirmamos

	return &product, nil
}

func (pr *ProductRepository) DeleteProductById(id_product int) error {
	query, err := pr.connection.Prepare("DELETE FROM products WHERE id_product=$1")

	if err != nil {
		fmt.Println(err)
		return err
	}

	defer query.Close()

	_, err = query.Exec(id_product)
	if err != nil {
		return err
	}
	return nil
}
