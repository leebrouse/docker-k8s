package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Auth Service: Authentication successful")
}

func main() {
	http.HandleFunc("/auth", handler)
	fmt.Println("Auth Service started on :8082")
	http.ListenAndServe(":8082", nil)
}
