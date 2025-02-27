package main

import "net/http"

func basicAuthHandler(w http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()
	if !ok || username != "alice@example.com" || password != "pa55word" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	w.Write([]byte("Authentication successful" + "\n"))
}
func main() {
	http.HandleFunc("/", basicAuthHandler)
	http.ListenAndServe(":8080", nil)
}
