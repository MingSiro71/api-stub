package main

import (
	"api_stub/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	controller := controllers.NewMainController()

	router.HandleFunc("/set/{id}", controller.Set).Methods("POST")
	router.HandleFunc("/get/{id}", controller.Get)
	router.HandleFunc("/list/{id}", controller.List).Methods("GET")
	router.HandleFunc("/clear/{id}", controller.Clear).Methods("PUT")
	router.HandleFunc("/init", controller.Init).Methods("PUT")

	router.HandleFunc("/", controller.Help).Methods("GET")

	// register router
	http.Handle("/", router)

	// start server with reverse proxy
	// http.ListenAndServe("localhost:8080", nil)
	//
	// start standalone server, expose 80
	http.ListenAndServe(":80", nil)
}
