package handlers

import (
	"encoding/json"
	"net/http"

	"019.restful_mongodb/models"
	"019.restful_mongodb/mongodb"

	"github.com/gorilla/mux"
)

var users []models.User

func GetUsers(w http.ResponseWriter, r *http.Request) {
	payload := mongodb.GetUsers()
	json.NewEncoder(w).Encode(payload)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	id := convertParams(r)
	payload := mongodb.GetUser(id)

	if payload.ID != "" {
		json.NewEncoder(w).Encode(payload)
		return
	}

	json.NewEncoder(w).Encode("User not found")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	mongodb.CreateUser(user)
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	id := convertParams(r)
	_ = json.NewDecoder(r.Body).Decode(&user)
	mongodb.UpdateUser(user, id)
	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := convertParams(r)
	mongodb.DeleteUser(id)
}

func convertParams(r *http.Request) string {
	params := mux.Vars(r)
	return params["id"]
}
