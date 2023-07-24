package controllers

import (
	"encoding/json"
	"net/http"
	"tc_go/db"
	u "tc_go/utils"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {

	account := &db.Account{}
	err := json.NewDecoder(r.Body).Decode(account) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := account.Create() //Create account
	u.Respond(w, resp)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {

	account := &db.Account{}
	err := json.NewDecoder(r.Body).Decode(account) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := db.Login(account.Email, account.Password)
	u.Respond(w, resp)
}
