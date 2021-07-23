package handlers

import (
	"log"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func (p Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.l.Printf("Logger")
}

func NewProducts(logger *log.Logger) *Products {
	return &Products{l: logger}
}
