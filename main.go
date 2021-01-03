package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/pawanpaudel93/go-mux-restapi/controller"
)

func main() {
	router := mux.NewRouter()

	controller.InitDatabase()

	router.HandleFunc("/resources", controller.GetResources).Methods("GET")
	router.HandleFunc("/resources/{id}", controller.GetResource).Methods("GET")
	router.HandleFunc("/resources", controller.CreateResource).Methods("POST")
	router.HandleFunc("/resources/{id}", controller.UpdateResource).Methods("PUT")
	router.HandleFunc("/resources/{id}", controller.DeleteResource).Methods("DELETE")

	srv := &http.Server{
		Handler: router,
		Addr:    ":5000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
