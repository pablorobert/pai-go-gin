package entity

type Product struct {
	ID    int     `json:"id_product"`
	Name  string  `json:"name_product"`
	Price float64 `json:"price_doruct"`
	Stock int     `json:"stock_product"`
	Code  int     `json:"code_product"`
}
