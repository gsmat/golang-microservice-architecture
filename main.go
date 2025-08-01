package main

import (
	"context"
	"golang-microservice-architecture/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	l := log.New(os.Stdout, "product-api ", log.LstdFlags)
	//helloHandler := handlers.NewHello(l)
	//goodByeHandler := handlers.NewGoodBye(l)

	productsHandler := handlers.NewProducts(l)

	serveMux := http.NewServeMux()
	//serveMux.Handle("/", helloHandler)
	//serveMux.Handle("/good-bye", goodByeHandler)
	serveMux.Handle("/products", productsHandler)

	log.Println("Starting server on :9090")

	server := http.Server{
		Addr:         ":9090",
		Handler:      serveMux,
		ErrorLog:     l,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, os.Interrupt)
	signal.Notify(signalChannel, os.Kill)

	sgnl := <-signalChannel
	log.Println("Got signal:", sgnl)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(tc)
}
