package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

type User struct {
	gorm.Model
	Name  string
	Email string
}

func InitialMigration() {
	db, err = gorm.Open("postgres", "postgres://klcxexdm:ROMQC6Po7yFUkAYK8NnSP8p1ejvb0g_a@ziggy.db.elephantsql.com:5432/klcxexdm")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}

	defer db.Close()

	db.AutoMigrate(&User{})
}

func AllUsers(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "All Users Endpoint Hit")
	db, err = gorm.Open("postgres", "postgres://klcxexdm:ROMQC6Po7yFUkAYK8NnSP8p1ejvb0g_a@ziggy.db.elephantsql.com:5432/klcxexdm")
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()

	var users []User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("postgres", "postgres://klcxexdm:ROMQC6Po7yFUkAYK8NnSP8p1ejvb0g_a@ziggy.db.elephantsql.com:5432/klcxexdm")
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	db.Create(&User{Name: name, Email: email})

	fmt.Fprint(w, "New User Successfully Created")

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("postgres", "postgres://klcxexdm:ROMQC6Po7yFUkAYK8NnSP8p1ejvb0g_a@ziggy.db.elephantsql.com:5432/klcxexdm")
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]

	var user User
	db.Where("name = ?", name).Find(&user)
	db.Delete(&user)

	fmt.Fprintf(w, "User Successfully Deleted")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update User Endpoint Hit")
}
