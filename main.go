package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/GoAKS", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!, It's Ethan Sumner.")
		fmt.Fprintf(w, "You Did it!")
	})
	fmt.Printf("Server running (port=8080), route: http://localhost:8080/GoAKS\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
