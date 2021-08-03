package data

import (
	"encoding/json"
	"io"
)

type Product struct {
	ID          int     `json:"id"` //have added field tags to the fields of the struct to modify the output json
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
}

// defining a type of []Product
type Products []*Product

// attaching a ToJSON method to the Products type
func (p *Products) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(GetProducts())
}

func (p *Product) FromJSON(reader io.Reader) error {
	e := json.NewDecoder(reader)
	return e.Decode(p)
}

func GetProducts() Products {
	return productList
}

func AddProduct(product *Product) Products {
	productList = append(productList, product)
	return productList
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Latte",
		Price:       5,
	},
}
