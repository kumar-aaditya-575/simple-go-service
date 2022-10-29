package main

import (
	"log"
	"net/http"
	"os"
	"trial/working/handlers"
)

func main() {
	logger := log.New(os.Stdout, "product-api", log.LstdFlags)
	helloHandler := handlers.NewHello(logger)
	goodbyeHandler := handlers.NewGoodbye(logger)
	serveMux := http.NewServeMux()
	serveMux.Handle("/hello", helloHandler)
	serveMux.Handle("/goodbye", goodbyeHandler)
	//http.HandleFunc("/hello", helloHandler)
	err := http.ListenAndServe(":9090", serveMux)
	if err != nil {
		return
	}

}
