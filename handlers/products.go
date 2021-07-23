package handlers

import (
	"GolangHttp/data"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

type Products struct {
	l *log.Logger
}

func (p Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProduct(w, r)
		return
	} else if r.Method == http.MethodPost {
		p.addProduct(w, r)
		return
	} else if r.Method == http.MethodPut {
		p.l.Printf("Path URL is %s", r.URL.Path)
		reg := regexp.MustCompile("/api/products/([0-9]+)")
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)
		if len(g) != 1 {
			http.Error(w, "Unable to Marshall JSON", http.StatusInternalServerError)
			return
		}

		idString := g[0][1]
		id, _ := strconv.Atoi(idString)

		fmt.Fprintf(w, "Id is %d", id)
		// TODO : Not implemented yet
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (p Products) getProduct(w http.ResponseWriter, r *http.Request) {
	prod := data.GetProducts()
	err := prod.ToJSON(w)

	if err != nil {
		http.Error(w, "Unable to Marshall JSON", http.StatusInternalServerError)
		return
	}

	p.l.Printf("Logger")
}

func (p Products) addProduct(w http.ResponseWriter, r *http.Request) {
	product := &data.Product{}
	err := product.FromJSON(r.Body)

	if err != nil {
		http.Error(w, "error", http.StatusBadRequest)
		return
	}

	response := data.AddProduct(product)
	resErr := response.ToJSON(w)

	if resErr != nil {
		http.Error(w, "error", http.StatusInternalServerError)
		return
	}

	p.l.Printf("Logger")
}

func NewProducts(logger *log.Logger) *Products {
	return &Products{l: logger}
}
