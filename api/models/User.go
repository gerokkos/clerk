package models

import (
	"time"
)

//Name struct
type Name struct {
	First string `json:"first"`
	Last  string `json:"last"`
}

//Registered struct
type Registered struct {
	Date time.Time `json:"date"`
}

//Picture struct
type Picture struct {
	Medium string `json:"medium"`
}

//User struct
type User struct {
	ID         uint       `json:"id"`
	Name       Name       `json:"name"`
	Email      string     `json:"email"`
	Cell       string     `json:"cell"`
	Picture    Picture    `json:"picture"`
	Registered Registered `json:"registered"`
}

//UserList struct
type UserList struct {
	Users []User `json:"results"`
}

//Filter struct
type Filter struct {
	Limit         int64  `schema:"limit"`
	Email         string `schema:"email"`
	StartingAfter int64  `schema:"starting_after"`
	EndingBefore  int64  `schema:"ending_before"`
}

//URL struct
type URL struct {
	URL string `schema:"url"`
}

//Clerks struct used to drop the table
type Clerks struct {
}
