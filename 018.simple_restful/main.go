package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type User struct {
	ID        int    `json:"id,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
	Info      `json:"info,omitempty"`
}
type Info struct {
	City  string `json:"city,omitempty"`
	Phone int    `json:"phone,omitempty"`
}

var users []User

func GetUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	id := convertParams(r)
	for _, u := range users {
		if u.ID == id {
			json.NewEncoder(w).Encode(u)
			return
		}
	}
	json.NewEncoder(w).Encode("User not found")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	users = append(users, user)
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	id := convertParams(r)
	_ = json.NewDecoder(r.Body).Decode(&user)
	for i, u := range users {
		if u.ID == id {
			users[i] = user
			json.NewEncoder(w).Encode(user)
			break
		}
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := convertParams(r)
	for i, u := range users {
		if u.ID == id {
			copy(users[i:], users[i+1:])
			users = users[:len(users)-1]
			break
		}
	}
	json.NewEncoder(w).Encode(users)
}

func convertParams(r *http.Request) int {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	return id
}

func main() {
	router := mux.NewRouter()
	user1 := User{ID: 1, Firstname: "hello", Lastname: "World", Info: Info{City: "Taipei", Phone: 123}}
	user2 := User{ID: 2, Firstname: "hello", Lastname: "World", Info: Info{City: "Taipei", Phone: 456}}
	users = append(users, user1, user2)
	router.HandleFunc("/users", GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", GetUser).Methods("GET")
	router.HandleFunc("/users", CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")
	fmt.Println("Starting server on port 3000...")
	log.Fatal(http.ListenAndServe(":3000", router))
}
