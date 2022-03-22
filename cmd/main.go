package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/robertamadge/veterinayClinic/pkg/handlers"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", handlers.GetTemplate)
	router.HandleFunc("/owners", handlers.GetAllOwners).Methods(http.MethodGet)
	router.HandleFunc("/owners/{id}", handlers.GetOwner).Methods(http.MethodGet)
	router.HandleFunc("/owners", handlers.CreateOwners).Methods(http.MethodPost)
	router.HandleFunc("/owners/{id}", handlers.UpdateOwner).Methods(http.MethodPut)
	router.HandleFunc("/owners/{id}", handlers.DeleteOwner).Methods(http.MethodDelete)
	fmt.Println(http.ListenAndServe(":8000", router))
}
