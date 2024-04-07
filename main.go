package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ahmedshakshak/books-recommender/register"
)

func main() {
	reg := register.NewRegister()
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), reg.App.Handler))
}
