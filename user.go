package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []User
	DB.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func GetUsersByName(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("firstname")
	last_name := r.URL.Query().Get("lastname")
	user_email := r.URL.Query().Get("email")

	if empty(name) {
		log.Println("Url Param 'name' is missing or empty")
		return
	}

	if empty(last_name) {
		last_name = ""

	}
	if empty(user_email) {
		user_email = ""
	}

	likeName := "%" + name + "%"
	likeLastName := "%" + last_name + "%"
	log.Println("Url Param 'key' is: " + likeName)
	w.Header().Set("Content-Type", "application/json")
	var users []User
	DB.Where("first_name LIKE ?  AND last_name LIKE ? AND email = ?", likeName, likeLastName, user_email).Table("users").Find(&users)

	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user User
	DB.First(&user, params["id"])
	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	DB.Create(&user)
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user User
	DB.First(&user, params["id"])
	json.NewDecoder(r.Body).Decode(&user)
	DB.Save(&user)
	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user User
	DB.Delete(&user, params["id"])
	json.NewEncoder(w).Encode("The USer is Deleted Successfully!")
}

func empty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}
