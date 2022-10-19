package handlers

import (
	"log"
	"net/http"

	"github.com/LarryCodes/products/models"
)

type ProductsHandler struct {
	l *log.Logger
}

func NewProductsHandler(l *log.Logger) *ProductsHandler {
	return &ProductsHandler{l}
}

func (h *ProductsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	productsList := models.GetProducts()
	err := productsList.ToJson(w)
	if err != nil {
		http.Error(w, "Unable to encode json", http.StatusBadRequest)
	}

}
