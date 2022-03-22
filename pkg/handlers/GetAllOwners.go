package handlers

import (
	"encoding/json"
	"github.com/robertamadge/veterinayClinic/pkg/mocks"
	"net/http"
)

func GetAllOwners(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mocks.Owners)
}