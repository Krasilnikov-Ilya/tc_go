package controllers

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"tc_go/db"
	u "tc_go/utils"
)

var GetHouseWithParkingPlacesAndPersonsById = func(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		panic(err)
	}
	data := db.GetHouseWithParkingPlacesAndPersons(id)
	resp := u.Message(data != nil, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var GetHousesWithParkingPlacesAndPersons = func(w http.ResponseWriter, r *http.Request) {
	data := db.GetHousesWithParkingPlacesAndPersons()
	resp := u.Message(data != nil, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
