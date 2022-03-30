package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/robertamadge/veterinayClinic/pkg/models"
	"net/http"
)

func (h handler) GetAllPets(w http.ResponseWriter, r *http.Request) {
	var pets []models.Pet

	if result := h.DB.Find(&pets); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pets)
}
