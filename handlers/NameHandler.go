package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Namehandler struct {
	logger *log.Logger
}

func (n Namehandler) ServeHTTP(writer http.ResponseWriter, request *http.Request)  {
	n.logger.Printf("Hello World")
	d,err := ioutil.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, "error", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(writer,"Hello World %s", d)
}

func NewNameHandler(l *log.Logger) *Namehandler {
	return &Namehandler{l}
}
