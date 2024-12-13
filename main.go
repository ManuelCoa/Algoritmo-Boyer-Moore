package main

import (
	"boyer-moore-service/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/search", handlers.SearchHandler)
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
