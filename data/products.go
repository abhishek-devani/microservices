package data

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"descriptioin"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"`
	CreatedOn   string  `json:"-"` // remove this field from response
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

type Products []*Product

func GetProducts() Products {
	return productList
}

func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

func UpdateProduct(p *Product, id int) error {

	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}

	p.ID = id
	productList[pos] = p

	return nil

}

var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (*Product, int, error) {
	for k, v := range productList {
		if v.ID == id {
			return v, k, nil
		}
	}
	return nil, 0, ErrProductNotFound
}

func getNextID() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

// for json marshal method
// func GetProducts() []*Product {
// 	return productList
// }

func (p *Product) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(p)
}

func (p *Product) Validate() error {

	validator := validator.New()

	validator.RegisterValidation("sku", validateSKU)

	return validator.Struct(p)

}

func validateSKU(fl validator.FieldLevel) bool {

	// format of sku is abc-aknds-askn

	re := regexp.MustCompile(`[a-z]+-[a-z}+-[a-z]+`)
	matches := re.FindAllString(fl.Field().String(), -1)

	log.Println(matches)

	return len(matches) == 1
}

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
