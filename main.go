package main

import (
	"log"
	"net/http"
	"phone-blocked-list-api/handlers"
)

func main() {
  http.HandleFunc("/", handlers.GetNumberHandler)
  http.HandleFunc("/add", handlers.AddNumberHandler)
  log.Println("Server is running at localhost:8080")
  http.ListenAndServe(":8080", nil)
}