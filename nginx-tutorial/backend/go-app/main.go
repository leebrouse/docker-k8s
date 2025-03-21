package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello from Go Web App!")
    })

    fmt.Println("Server is running on port 8081...")
    http.ListenAndServe(":8081", nil)
}