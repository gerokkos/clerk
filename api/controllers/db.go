package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gerokkos/clerk/api/models"
)

func (server *Server) seed(url string) error {
	db := OpenConnection()
	if url == "" {
		url = "https://randomuser.me/api/?results=5000&inc=name,email,cell,registered,picture&noinfo"
	}
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	var results models.UserList

	json.Unmarshal([]byte(bodyBytes), &results)

	sqlStatement := `
	INSERT INTO clerks (first_name, last_name, email, cell, picture, registered_on)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING user_id`

	for k := range results.Users {
		_, err = db.Exec(sqlStatement, results.Users[k].Name.First, results.Users[k].Name.Last, results.Users[k].Email, results.Users[k].Cell, results.Users[k].Picture.Medium, results.Users[k].Registered.Date)
		if err != nil {
			panic(err)
		} else {
		}
	}
	defer db.Close()
	return err
}

func (server *Server) getAllUsers(limit int64, email string, startingAfter int64, endingBefore int64) ([]models.User, error) {
	// create the postgres db connection
	db := OpenConnection()
	defer db.Close()
	var users []models.User

	if limit == 0 {
		limit = 10
	}

	q := fmt.Sprintf("SELECT * FROM clerks")
	args := []interface{}{}

	// Add conditional query/args
	if email != "" || startingAfter != 0 || endingBefore != 0 {
		q = fmt.Sprintf("%s WHERE ", q)
		args = append(args)
	}
	if email != "" && startingAfter == 0 && endingBefore == 0 {
		q = fmt.Sprintf("%s email=$1", q)
		args = append(args, email)
	}
	if startingAfter != 0 {
		q = fmt.Sprintf("%s user_id > $1 AND user_id BETWEEN $2 and $3 ORDER BY registered_on DESC LIMIT $4", q)
		args = append(args, startingAfter, startingAfter, startingAfter+limit, limit)
	}
	if endingBefore != 0 {
		q = fmt.Sprintf("%s user_id < $1 AND user_id BETWEEN $2 and $3 ORDER BY registered_on DESC LIMIT $4", q)
		args = append(args, endingBefore, endingBefore-limit, endingBefore, limit)
	}
	if email == "" && startingAfter == 0 && endingBefore == 0 {
		q = fmt.Sprintf("%s ORDER BY registered_on DESC LIMIT $1", q)
		args = append(args, limit)
	}
	if limit == 0 || limit > 100 {
		log.Fatalf("limit should be from 1 to 100")
	}

	rows, err := db.Query(q, args...)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.ID, &user.Name.First, &user.Name.Last, &user.Email, &user.Cell, &user.Picture.Medium, &user.Registered.Date)
		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}
		users = append(users, user)
	}
	return users, err
}
