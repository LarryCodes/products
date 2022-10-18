package main

import (
	"log"
	"net/http"
	"time"

	"github.com/LarryCodes/products/handlers"
)

func main() {

	productsHandler := handlers.NewProductsHandler()

	mux := http.NewServeMux()
	mux.Handle("/", productsHandler)

	server := &http.Server{
		Addr:         ":9000",
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
		IdleTimeout:  time.Second * 120,
	}

	log.Fatalln(server.ListenAndServe())

}
