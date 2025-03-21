package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Log Service: Received request")
}

func main() {
	http.HandleFunc("/log", handler)
	fmt.Println("Log Service started on :8083")
	http.ListenAndServe(":8083", nil)
}
