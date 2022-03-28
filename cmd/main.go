package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/robertamadge/veterinayClinic/pkg/db"
	"github.com/robertamadge/veterinayClinic/pkg/handlers"
	"net/http"
)


func main() {
	DB := db.Init()
	h := handlers.New(DB)
	router := mux.NewRouter()

	router.HandleFunc("/", h.GetTemplate)

	router.HandleFunc("/owners", h.GetAllOwners).Methods(http.MethodGet)
	router.HandleFunc("/owners/{id}", h.GetOwner).Methods(http.MethodGet)
	router.HandleFunc("/owners", h.CreateOwners).Methods(http.MethodPost)
	router.HandleFunc("/owners/{id}", h.UpdateOwner).Methods(http.MethodPut)
	router.HandleFunc("/owners/{id}", h.DeleteOwner).Methods(http.MethodDelete)
	fmt.Println(http.ListenAndServe(":8000", router))
}

