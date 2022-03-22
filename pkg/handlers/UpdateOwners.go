package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/robertamadge/veterinayClinic/pkg/mocks"
	"github.com/robertamadge/veterinayClinic/pkg/models"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func UpdateOwner(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var updateOwner models.Owner
	json.Unmarshal(body, &updateOwner)

	for index, owner := range mocks.Owners {
		if owner.Id == id {
			owner.Cpf = updateOwner.Cpf
			owner.Name = updateOwner.Name
			owner.Email = updateOwner.Email
			owner.Mobile = updateOwner.Mobile
			owner.Pets = updateOwner.Pets

			mocks.Owners[index] = owner

			w.WriteHeader(http.StatusOK)
			w.Header().Add("Content-type", "application/json")
			json.NewEncoder(w).Encode("Updated")
			break
		}
	}
}
