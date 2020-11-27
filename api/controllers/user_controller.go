package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gerokkos/clerk/api/models"
)

func (server *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	db := OpenConnection()
	resp, err := http.Get("https://randomuser.me/api/?results=5000&inc=name,email,cell,registered,picture&noinfo")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	var results models.UserList

	json.Unmarshal([]byte(bodyBytes), &results)

	sqlStatement := `
	INSERT INTO clerk (first_name, last_name, email, cell, picture, registered_on)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING user_id`

	for k := range results.Users {
		_, err = db.Exec(sqlStatement, results.Users[k].Name.First, results.Users[k].Name.Last, results.Users[k].Email, results.Users[k].Cell, results.Users[k].Picture.Medium, results.Users[k].Registered.Date)
		if err != nil {
			panic(err)
		} else {
		}
	}
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	defer db.Close()
}
