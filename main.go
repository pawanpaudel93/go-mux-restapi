package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/pawanpaudel93/go-mux-restapi/controller"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	router := mux.NewRouter()

	controller.InitDatabase()

	router.HandleFunc("/books", controller.GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", controller.GetBook).Methods("GET")
	router.HandleFunc("/books", controller.CreateBook).Methods("POST")
	router.HandleFunc("/books/{id}", controller.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", controller.DeleteBook).Methods("DELETE")

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
