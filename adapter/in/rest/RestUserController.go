package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)



var CreateUser = func(w http.ResponseWriter, r *http.Request) {
	user := &CreatedUserDto{}
	reqBody, err := ioutil.ReadAll(r.Body)
	log.Printf(string(reqBody))
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}
	json.Unmarshal(reqBody, user)
	userDomain:=user.ToDomain()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userDomain.Id)
}