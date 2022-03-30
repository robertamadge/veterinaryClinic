package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/robertamadge/veterinayClinic/pkg/models"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func (h handler) UpdateAppointment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var updateAppointment models.Appointment
	json.Unmarshal(body, &updateAppointment)

	var appointment models.Appointment
	if result := h.DB.First(&appointment, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	appointment.Date = updateAppointment.Date

	h.DB.Save(&appointment)

	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Updated")

}
