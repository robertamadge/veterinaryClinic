package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/robertamadge/veterinayClinic/pkg/models"
	"net/http"
)

func (h handler) GetAllAppointments(w http.ResponseWriter, r *http.Request) {
	var appointments []models.Appointment

	if result := h.DB.Find(&appointments); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(appointments)
}
