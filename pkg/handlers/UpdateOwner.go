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

func (h handler) UpdateOwner(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var updateOwner models.Owner
	json.Unmarshal(body, &updateOwner)

	var owner models.Owner
	if result := h.DB.First(&owner, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	owner.Cpf = updateOwner.Cpf
	owner.Name = updateOwner.Name
	owner.Email = updateOwner.Email
	owner.Username = updateOwner.Username
	owner.Mobile = updateOwner.Mobile

	h.DB.Save(&owner)

	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Updated")

}
