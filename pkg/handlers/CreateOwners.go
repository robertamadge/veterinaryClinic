package handlers

import (
	"encoding/json"
	"github.com/robertamadge/veterinayClinic/pkg/mocks"
	"github.com/robertamadge/veterinayClinic/pkg/models"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
)

func CreateOwners(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var owner models.Owner
	json.Unmarshal(body, &owner)

	owner.Id = rand.Intn(100)
	mocks.Owners = append(mocks.Owners, owner)

	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode("Created")

}