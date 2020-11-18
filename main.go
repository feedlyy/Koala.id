package main

import (
	"Koala/Rest/app/controllers"
	"net/http"
)

func main () {
	//list of routes
	http.HandleFunc("/register", controllers.Register)
	http.HandleFunc("/getToken", controllers.GetToken)

	//serve a server
	_ = http.ListenAndServe(":8000", nil)
}
