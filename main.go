package main

import (
	"fmt"
	"log"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	title := "LearTech Preview Certs work or do they !!!!"

	from := ""
	if r.URL != nil {
		from = r.URL.String()
	}
	if from != "/favicon.ico" {
		log.Printf("title: %s\n", title)
	}

	fmt.Fprintf(w, "Hello from : %s\n", title)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	title := "Health Check"

	from := ""
	if r.URL != nil {
		from = r.URL.String()
	}
	if from != "/favicon.ico" {
		log.Printf("title: %s\n", title)
	}

	fmt.Fprintf(w, "Healthy!")
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/health", healthHandler)
	http.ListenAndServe(":8080", nil)
}
