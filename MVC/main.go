package main

import (
	"mvc/controller"
	"net/http"
)

func main() {

	http.HandleFunc("/", controller.UserShowHandler)

	http.ListenAndServe(":8080", nil)

}
