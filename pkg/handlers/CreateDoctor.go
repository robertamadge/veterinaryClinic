package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/robertamadge/veterinayClinic/pkg/models"
	"io/ioutil"
	"log"
	"net/http"
)

func (h handler) CreateDoctor(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var doctor models.Doctor
	json.Unmarshal(body, &doctor)

	// Append to the Doctor table
	if result := h.DB.Create(&doctor); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")

}
