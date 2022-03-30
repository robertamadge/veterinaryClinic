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

func (h handler) UpdatePet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var updatePet models.Pet
	json.Unmarshal(body, &updatePet)

	var pet models.Pet
	if result := h.DB.First(&pet, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	pet.Name = updatePet.Name
	pet.Age = updatePet.Age

	h.DB.Save(&pet)

	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Updated")

}
