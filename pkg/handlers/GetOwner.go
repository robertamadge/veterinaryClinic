package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/robertamadge/veterinayClinic/pkg/mocks"
	"net/http"
	"strconv"
)

func GetOwner(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	for _ , owner := range mocks.Owners {
		if owner.Id == id {
			w.WriteHeader(http.StatusOK)
			w.Header().Add("Content-type", "application/json")
			json.NewEncoder(w).Encode(owner)
			break
		}
	}
}
