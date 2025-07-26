package handlers

import (
	"log"
	"net/http"
)

type GoodBye struct {
	log *log.Logger
}

func NewGoodBye(log *log.Logger) *GoodBye {
	return &GoodBye{log}
}

func (g *GoodBye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	g.log.Println("GoodBye World User")
	rw.Write([]byte("GoodBye World User"))
}
