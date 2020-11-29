package controllers

import (
	"log"

	"github.com/gerokkos/clerk/api/models"
	"github.com/jinzhu/gorm"
)

//Load is the function to drop and create the Clerks table
func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Clerks{}, &models.Clerks{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}

	err = db.Exec(`CREATE TABLE clerks (
		user_id serial PRIMARY KEY,
		first_name varchar(255) NOT NULL,
		last_name varchar(255) NOT NULL,
		email varchar(255),
		cell varchar(255),
		picture varchar(255),
		registered_on date
	)`).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}
}
