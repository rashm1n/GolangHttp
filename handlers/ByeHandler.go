package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type ByeHandler struct {
	logger *log.Logger
}

func (n ByeHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request)  {
	n.logger.Printf("Bye World")
	d,err := ioutil.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, "error", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(writer,"Bye World %s", d)
}

func NewByeHandler(l *log.Logger) *ByeHandler {
	return &ByeHandler{l}
}
