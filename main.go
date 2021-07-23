package main

import (
	"GolangHttp/handlers"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	logger := log.New(os.Stdout, "api", log.LstdFlags)

	productsHandler := handlers.NewProducts(logger)

	productServeMux := http.NewServeMux()
	productServeMux.Handle("/api/products", productsHandler)
	// create a new http server
	server := http.Server{
		Addr:         ":9090",
		Handler:      productServeMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// starts a server and binds the DefaultServeMux
	//http.ListenAndServe(":9090",sm)
	server.ListenAndServe()
}
