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

func (h handler) UpdateDoctor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var updateDoctor models.Doctor
	json.Unmarshal(body, &updateDoctor)

	var doctor models.Doctor
	if result := h.DB.First(&doctor, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	doctor.Crm = updateDoctor.Crm
	doctor.Cpf = updateDoctor.Cpf
	doctor.Name = updateDoctor.Name

	h.DB.Save(&doctor)

	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Updated")

}
