package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is Stationery maneger.Hello, World!")
}

type Stationery struct {
	ID        string    `json:"id"`
	Category  string    `json:"category"`
	Name      string    `json:"name"`
	Location  string    `json:"location"`
	InUse     bool      `json:"in_use"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func main() {
	fmt.Println("Starting HTTP server on http://localhost:8080")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
