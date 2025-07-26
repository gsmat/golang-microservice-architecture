package handlers

import (
	"encoding/json"
	"golang-microservice-architecture/data"
	"log"
	"net/http"
)

type Products struct {
	log *log.Logger
}

func NewProducts(log *log.Logger) *Products {
	return &Products{log}
}

func (products *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	listOfProducts := data.GetAllProducts()
	productData, err := json.Marshal(listOfProducts)
	if err != nil {
		http.Error(rw, "Unable to read products Json", http.StatusInternalServerError)
	}
	rw.Write(productData)
}
