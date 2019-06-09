package controllers

import (
	"encoding/json"
	"net/http"
	"preposterous/models"
	u "preposterous/utils"
	"strconv"

	"github.com/gorilla/mux"
)

var CreateVerb = func(w http.ResponseWriter, r *http.Request) {

	verb := &models.Verb{}

	err := json.NewDecoder(r.Body).Decode(verb)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}
	resp := verb.Create()
	u.Respond(w, resp)
}

var GetVerb = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		u.Respond(w, u.Message(false, "There was an error responding to your request"))
		return
	}

	data := models.GetVerb(uint(id))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
