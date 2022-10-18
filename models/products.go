package models

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
	Sku         string
	CreatedAt   string
	UpdatedAt   string
	DeletedAt   string
}

type Products []*Product

var productsList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "From the milky coffee",
		Price:       2.45,
		Sku:         "abc123",
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Esspresso",
		Description: "Short and strong without coffee",
		Price:       1.99,
		Sku:         "7845y8",
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
	},
}

func GetProducts() Products {
	return productsList
}

func (p *Products) ToJson(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(p)
}
