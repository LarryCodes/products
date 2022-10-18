package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
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

	go func() {
		log.Println("Starting server...")
		err := server.ListenAndServe()
		if err != nil {
			log.Println("Error: ", err)
		}
	}()

	// Use channel to bloack until interrupt is received and gracefully shutdown
	signalChannel := make(chan os.Signal)

	signal.Notify(signalChannel, os.Kill)
	signal.Notify(signalChannel, os.Interrupt)

	// Block until signal is received
	msg := <-signalChannel

	log.Println("Received", msg, "signal, gracefully shutting down!")

	timeOutContext, cancel := context.WithTimeout(context.Background(), time.Second*120)
	defer cancel()
	server.Shutdown(timeOutContext)

}
