package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/robertamadge/veterinayClinic/pkg/mocks"
	"net/http"
	"strconv"
)

func DeleteOwner(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	for index, owner := range mocks.Owners {
		if owner.Id == id {
			//delete owner and send response if the owner Id matches dynamic Id
			mocks.Owners = append(mocks.Owners[:index], mocks.Owners[index+1:]...)

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Deleted")
			break
		}
	}
}