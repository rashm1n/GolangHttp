package handlers

import (
	"fmt"
	"net/http"
)

type RootHandler struct {

}

func (r RootHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w,"Hello World\n")
}

func NewRootHandler() RootHandler {
	return RootHandler{}
}