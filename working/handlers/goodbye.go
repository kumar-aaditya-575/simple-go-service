package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type GoodBye struct {
	logger *log.Logger
}

func NewGoodbye(logger *log.Logger) *GoodBye {
	return &GoodBye{logger}
}
func (goodBye *GoodBye) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	goodBye.logger.Println("Goodbye World")
	body, err := io.ReadAll(request.Body)
	if err != nil {
		http.Error(responseWriter, "Oops", http.StatusBadRequest)
		return
	}
	log.Printf("Request Body: %s \n", body)
	fmt.Fprintf(responseWriter, "Goodbye %s", body)
}
