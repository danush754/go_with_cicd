package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	const port = ":8080"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Happy new year every one !")
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "The server runs fine...")
	})

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("failed to create a server: %v", err.Error())
	}
}
