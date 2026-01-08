package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is Stationery maneger.Hello, World!")
}

func main() {
	fmt.Println("Starting HTTP server on http://localhost:8080")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
