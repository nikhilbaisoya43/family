package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DNS = "root:root@123@tcp(127.0.0.1:3306)/employeedb?charset=utf8mb4&parseTime=True&loc=Local"

func initialMigration() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
	DB.AutoMigrate(&User{}, &FamilyType{})
}

func initializeRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/users", GetUsers).Methods("GET")
	r.HandleFunc("/search/users", GetUsersByName).Methods("GET")
	r.HandleFunc("/users/{id}", GetUser).Methods("GET")
	r.HandleFunc("/users", CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")

	r.HandleFunc("/family", GetAllFamily).Methods("GET")
	r.HandleFunc("/family", CreateFamily).Methods("POST")
	r.HandleFunc("/family/{id}", UpdateFamily).Methods("PUT")
	r.HandleFunc("/family/{id}", DeleteFamily).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9000", r))
	log.Println("server has started successfully")
}

func main() {
	initialMigration()
	initializeRouter()

}
