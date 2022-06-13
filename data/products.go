package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"descriptioin"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"` // remove this field from response
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

type Products []*Product

func GetProducts() Products {
	return productList
}

// for json marshal method
// func GetProducts() []*Product {
// 	return productList
// }

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w).Encode(p)
	return e
}

var productList = []*Product{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy Milky Coffee",
		Price:       2.45,
		SKU:         "abc123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
		DeletedOn:   "",
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Sort and strong coffee without milk",
		Price:       1.99,
		SKU:         "abc456",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
		DeletedOn:   "",
	},
}
