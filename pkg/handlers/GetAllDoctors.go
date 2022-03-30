package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/robertamadge/veterinayClinic/pkg/models"
	"net/http"
)

func (h handler) GetAllDoctors(w http.ResponseWriter, r *http.Request) {
	var doctors []models.Doctor

	if result := h.DB.Find(&doctors); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(doctors)
}
