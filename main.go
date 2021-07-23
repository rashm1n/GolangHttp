package main

import (
	"GolangHttp/handlers"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	logger := log.New(os.Stdout, "api",log.LstdFlags)
	nameHandler := handlers.NewNameHandler(logger)
	byeHandler := handlers.NewByeHandler(logger)
	rootHandler := handlers.NewRootHandler()

	sm := http.NewServeMux()
	sm.Handle("/name",nameHandler)
	sm.Handle("/bye",byeHandler)
	sm.Handle("/",rootHandler)

	// create a new http server
	server := http.Server{
		Addr:    ":9090",
		Handler: sm,
		IdleTimeout: 120*time.Second,
		ReadTimeout: 1*time.Second,
		WriteTimeout: 1*time.Second,
	}
	
	// starts a server and binds the DefaultServeMux
	//http.ListenAndServe(":9090",sm)
	server.ListenAndServe()
}