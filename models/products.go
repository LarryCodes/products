package models

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Sku         string  `json:"sku"`
	CreatedAt   string  `json:"-"`
	UpdatedAt   string  `json:"-"`
	DeletedAt   string  `json:"-"`
}

type Products []*Product

var productsList = []*Product{
	{
		ID:          1,
		Name:        "Latte",
		Description: "From the milky coffee",
		Price:       2.45,
		Sku:         "abc123",
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Esspresso",
		Description: "Short and strong without coffee",
		Price:       1.99,
		Sku:         "7845y8",
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
	},
}

func (p *Product) FromJson(r *http.Request) error {
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(p)
}

func AddProduct(p *Product) {
	// First add the correct id
	p.ID = getNextId()
	//Append the product to the products list slice
	productsList = append(productsList, p)
}

func GetProducts() Products {
	return productsList
}

func (p *Products) ToJson(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(p)
}

func getNextId() int {
	lastIdx := productsList[len(productsList)-1]

	return lastIdx.ID + 1
}
