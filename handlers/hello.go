package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	log *log.Logger
}

func NewHello(log *log.Logger) *Hello {
	return &Hello{log}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.log.Println("Handle Hello Request")

	//read the body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.log.Println("Error reading body:", err)
		http.Error(rw, "Unable to read request Body", http.StatusInternalServerError)
		return
	}

	//write the response
	fmt.Fprintf(rw, "Hello %s\n", body)
}
