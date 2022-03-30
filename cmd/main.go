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
	router.HandleFunc("/owners", h.CreateOwner).Methods(http.MethodPost)
	router.HandleFunc("/owners/{id}", h.UpdateOwner).Methods(http.MethodPut)
	router.HandleFunc("/owners/{id}", h.DeleteOwner).Methods(http.MethodDelete)

	router.HandleFunc("/pets", h.GetAllPets).Methods(http.MethodGet)
	router.HandleFunc("/pets/{id}", h.GetPet).Methods(http.MethodGet)
	router.HandleFunc("/pets", h.CreatePet).Methods(http.MethodPost)
	router.HandleFunc("/pets/{id}", h.UpdatePet).Methods(http.MethodPut)
	router.HandleFunc("/pets/{id}", h.DeletePet).Methods(http.MethodDelete)

	router.HandleFunc("/doctors", h.GetAllDoctors).Methods(http.MethodGet)
	router.HandleFunc("/doctors/{id}", h.GetDoctor).Methods(http.MethodGet)
	router.HandleFunc("/doctors", h.CreateDoctor).Methods(http.MethodPost)
	router.HandleFunc("/doctors/{id}", h.UpdateDoctor).Methods(http.MethodPut)
	router.HandleFunc("/doctors/{id}", h.DeleteDoctor).Methods(http.MethodDelete)

	router.HandleFunc("/appointments", h.GetAllAppointments).Methods(http.MethodGet)
	router.HandleFunc("/appointments/{id}", h.GetAppointment).Methods(http.MethodGet)
	router.HandleFunc("/appointments", h.CreateAppointment).Methods(http.MethodPost)
	router.HandleFunc("/appointments/{id}", h.UpdateAppointment).Methods(http.MethodPut)
	router.HandleFunc("/appointments/{id}", h.DeleteAppointment).Methods(http.MethodDelete)
	fmt.Println(http.ListenAndServe(":8000", router))
}

