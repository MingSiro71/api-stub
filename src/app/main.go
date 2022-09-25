package main

import (
	"api_stub/controllers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	router := mux.NewRouter()
	controller := controllers.NewMainController()
	
	htmlPath := filepath.Join(os.Getenv("APPROOT"), "html");
	publicPath := filepath.Join(os.Getenv("APPROOT"), "public");

	router.HandleFunc("/set/{id}", controller.Set).Methods("POST")
	router.HandleFunc("/get/{id}", controller.Get)
	router.HandleFunc("/list/{id}", controller.List).Methods("GET")
	router.HandleFunc("/clear/{id}", controller.Clear).Methods("PUT")
	router.HandleFunc("/init", controller.Init).Methods("PUT")
	router.Handle("/", http.FileServer(http.Dir(htmlPath)))

	// static roure
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(publicPath))))

	// register router
	http.Handle("/", router)

	// start server with reverse proxy
	// http.ListenAndServe("localhost:8080", nil)
	//
	// start standalone server, expose 80
	http.ListenAndServe(":80", nil)
}
