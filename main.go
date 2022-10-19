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

	logger := log.New(os.Stdout, "Products-api:", log.LstdFlags)

	productsHandler := handlers.NewProductsHandler(logger)

	mux := http.NewServeMux()
	mux.Handle("/", productsHandler)

	server := &http.Server{
		Addr:         ":9000",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// Run the server in a seperate goroutine
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Println(err)
		}
	}()

	// Use channels to block
	sigChan := make(chan os.Signal)

	// Notify the channel on os.Kill or os.Interrupt
	signal.Notify(sigChan, os.Kill)
	signal.Notify(sigChan, os.Interrupt)

	sigData := <-sigChan

	// Gracefully shutdown
	logger.Println("Received", sigData, "Signal, gracefully shutting down.")

	timeOutContext, cancelFunc := context.WithTimeout(context.Background(), time.Second*30)
	defer cancelFunc()

	server.Shutdown(timeOutContext)

}
