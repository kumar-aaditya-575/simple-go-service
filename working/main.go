package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
)

func main() {
	fmt.Println(math.Abs(2.4 / 1.2))

	//router := mux.BuildVarsFunc()

	http.HandleFunc("/", func(responseWriter http.ResponseWriter, request *http.Request) {
		log.Println("Hello World")
	})
	http.ListenAndServe(":9090", nil)

}
