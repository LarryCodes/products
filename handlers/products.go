package handlers

import (
	"net/http"

	"github.com/LarryCodes/products/models"
)

type ProductsHandler struct{}

func NewProductsHandler() *ProductsHandler {
	return &ProductsHandler{}
}

func (h *ProductsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	productsList := models.GetProducts()
	err := productsList.ToJson(w)
	if err != nil {
		http.Error(w, "Unable to encode json", http.StatusBadRequest)
	}

}
