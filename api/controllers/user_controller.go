package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gerokkos/clerk/api/models"
	"github.com/gerokkos/clerk/api/responses"
	"github.com/gorilla/schema"
)

//Populate is used to call the Randomuser Api to seed the DB
func (server *Server) Populate(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Print("Fail to parse!")
	}

	url := new(models.URL)
	if err := schema.NewDecoder().Decode(url, r.Form); err != nil {
		fmt.Print("Fail to decode!")
	}
	fmt.Printf("%+v", url)
	err := server.seed(string(url.URL))

	if err == nil {
		json.NewEncoder(w).Encode("Success!")
	}
}

//Clerks is used from the GET endpoint
func (server *Server) Clerks(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Print("Fail to parse!")
	}

	filter := new(models.Filter)
	if err := schema.NewDecoder().Decode(filter, r.Form); err != nil {
		fmt.Print("Fail to decode!")
	}

	fmt.Printf("%+v", filter)
	email := strings.ToLower(filter.Email)
	users, err := server.getAllUsers(int64(filter.Limit), string(email), int64(filter.StartingAfter), int64(filter.EndingBefore))

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, users)
}
