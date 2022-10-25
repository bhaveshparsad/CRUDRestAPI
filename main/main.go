package main

import (
	"fmt"
	"log"
	"net/http"

	"CRUDRestAPI/controllers"

	"github.com/gorilla/mux"
)

func route() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/GetAllContacts", service.GetAllContacts).Methods("GET")
	r.HandleFunc("/GetContact/{id}", service.GetContact).Methods("GET")
	r.HandleFunc("/CreateContact", service.CreateContact).Methods("POST")
	r.HandleFunc("/UpdateContact/{id}", service.UpdateContact).Methods("PUT")
	r.HandleFunc("/DeleteContact", service.DeleteContact).Methods("DELETE")

	return r
}

func main() {

	r := route()

	log.Fatal(http.ListenAndServe(":12345", r))

	fmt.Println("Success...")
}
