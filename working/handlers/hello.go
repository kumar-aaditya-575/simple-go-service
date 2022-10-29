package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Hello struct {
	logger *log.Logger
}

func NewHello(logger *log.Logger) *Hello {
	return &Hello{logger}
}

func (hello *Hello) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	hello.logger.Println("Hello World")
	body, err := io.ReadAll(request.Body)
	if err != nil {
		http.Error(responseWriter, "Oops", http.StatusBadRequest)
		return
	}
	log.Printf("Request Body: %s \n", body)
	fmt.Fprintf(responseWriter, "Hello %s", body)
}
