package controllers

import (
	f "fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"tc_go/db"
	u "tc_go/utils"
)

var GetPersonById = func(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		panic(err)
	}
	data := db.GetPerson(id)
	resp := u.Message(data != nil, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var GetPersons = func(w http.ResponseWriter, r *http.Request) {
	data := db.GetPersons()
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
	f.Println()
}
