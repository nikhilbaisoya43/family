package main

import (
	"encoding/json"

	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type FamilyType struct {
	gorm.Model
	Name     string `json:"name"`
	Relation string `json:"Relation"`
}

func GetAllFamily(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var family []FamilyType
	DB.Find(&family)
	json.NewEncoder(w).Encode(family)
}

func CreateFamily(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var family FamilyType
	json.NewDecoder(r.Body).Decode(&family)
	DB.Create(&family)
	json.NewEncoder(w).Encode(family)
}

func UpdateFamily(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	var family FamilyType
	DB.First(&family, params["id"])
	json.NewDecoder(r.Body).Decode(&family)
	DB.Save(&family)
	json.NewEncoder(w).Encode(family)
}

func DeleteFamily(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	var family FamilyType
	DB.Delete(&family, params["id"])
	json.NewEncoder(w).Encode("The family is Deleted Successfully!")

}
