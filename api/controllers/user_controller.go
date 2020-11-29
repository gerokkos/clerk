package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gerokkos/clerk/api/models"
	"github.com/gorilla/schema"
)

func (server *Server) Populate(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		// Handle error
	}

	url := new(models.Url)
	if err := schema.NewDecoder().Decode(url, r.Form); err != nil {
		// Handle error
	}
	// Do something with filter
	fmt.Printf("%+v", url)
	err := server.seed(string(url.Url))

	if err == nil {
		json.NewEncoder(w).Encode("Success!")
	}
}

func (server *Server) Clerks(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		// Handle error
	}

	filter := new(models.Filter)
	if err := schema.NewDecoder().Decode(filter, r.Form); err != nil {
		// Handle error
	}
	// Do something with filter
	fmt.Printf("%+v", filter)
	email := strings.ToLower(filter.Email)
	users, err := server.getAllUsers(int64(filter.Limit), string(email), int64(filter.StartingAfter), int64(filter.EndingBefore))

	if err != nil {
		log.Fatalf("Unable to get clerks. %v", err)
	}
	// send all the users as response
	json.NewEncoder(w).Encode(users)
}
