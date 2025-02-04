package main

import (
	"fmt"
	"go-url-shortener/internal/routes"
	"net/http"
)

func main() {
	router := routes.NewRouter()

	fmt.Println("Server is listening")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}
