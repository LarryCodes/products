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

	if r.Method == http.MethodGet {
		h.getProducts(w)
		return
	}

	if r.Method == http.MethodPost {
		h.addProduct(w, r)
		return
	}

}

func (p *ProductsHandler) getProducts(w http.ResponseWriter) {
	productsList := models.GetProducts()
	err := productsList.ToJson(w)
	if err != nil {
		http.Error(w, "Unable to encode json", http.StatusBadRequest)
	}
}

func (p *ProductsHandler) addProduct(w http.ResponseWriter, r *http.Request) {

	data := &models.Product{}
	err := data.FromJson(r)
	if err != nil {
		p.l.Println(err)
		http.Error(w, "Unable to decode json object", http.StatusBadRequest)
	}

	models.AddProduct(data)
}
