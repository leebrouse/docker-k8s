package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Payment Service: Payment processed successfully")
}

func main() {
	http.HandleFunc("/payment", handler)
	fmt.Println("Payment Service started on :8084")
	http.ListenAndServe(":8084", nil)
}
