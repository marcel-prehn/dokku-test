package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting Service on Port 8080")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Golang running on Dokku!")
	})

	http.ListenAndServe(":8080", nil)
}
