package main

import (
	"fmt"
	"log"
	"net/http"

	"api/handlers"
)

func main() {
	http.HandleFunc("/view-data", handlers.ReadFile)
	http.HandleFunc("/submit", handlers.WriteFile)

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
