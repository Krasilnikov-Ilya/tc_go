package controllers

import (
	f "fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"tc_go/db"
	u "tc_go/utils"
)

var GetCarById = func(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		panic(err)
	}
	data := db.GetCar(id)
	resp := u.Message(data != nil, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var GetCars = func(w http.ResponseWriter, r *http.Request) {
	data := db.GetCars()
	resp := u.Message(data != nil, "success")
	resp["data"] = data
	u.Respond(w, resp)
	f.Println()
}

var GetPersonsCarsById = func(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		panic(err)
	}
	data := db.GetPersonsCarsById(id)
	resp := u.Message(data != nil, "success")
	resp["data"] = data
	u.Respond(w, resp)
	f.Println()
}
