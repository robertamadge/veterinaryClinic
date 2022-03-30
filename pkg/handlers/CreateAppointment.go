package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/robertamadge/veterinayClinic/pkg/models"
	"io/ioutil"
	"log"
	"net/http"
)

func (h handler) CreateAppointment(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var appointment models.Appointment
	json.Unmarshal(body, &appointment)

	// Append to the Appointment table
	if result := h.DB.Create(&appointment); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")

}
