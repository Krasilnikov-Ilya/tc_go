package controllers

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"tc_go/db"
	u "tc_go/utils"
)

var GetPersonWithCarById = func(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		panic(err)
	}
	personWithCars := db.GetPersonWithCars(id)
	resp := u.Message(personWithCars != nil, "success")
	resp["data"] = personWithCars
	u.Respond(w, resp)
}
