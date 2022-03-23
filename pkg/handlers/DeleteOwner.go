package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/robertamadge/veterinayClinic/pkg/models"
	"net/http"
	"strconv"
)

func (h handler) DeleteOwner(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var owner models.Owner

	if result := h.DB.First(&owner, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	//Delete that owner
	h.DB.Delete(&owner)

	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted")

}
