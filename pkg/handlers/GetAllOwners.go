package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/robertamadge/veterinayClinic/pkg/models"
	"net/http"
)

func (h handler) GetAllOwners(w http.ResponseWriter, r *http.Request) {
	var owners []models.Owner

	if result := h.DB.Find(&owners); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(owners)
}