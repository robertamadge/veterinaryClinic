package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/robertamadge/veterinayClinic/pkg/models"
	"net/http"
	"strconv"
)

func (h handler) DeleteAppointment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var appointment models.Appointment

	if result := h.DB.First(&appointment, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	//Delete that owner
	h.DB.Delete(&appointment)

	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted")

}
