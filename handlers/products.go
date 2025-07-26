package handlers

import (
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
	if r.Method == http.MethodGet {
		products.GetProducts(rw, r)
		return
	}

	//catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (products *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	listOfProducts := data.GetAllProducts()
	err := listOfProducts.ToJson(rw)
	if err != nil {
		http.Error(rw, "Unable to read products Json", http.StatusInternalServerError)
	}
}
