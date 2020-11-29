package controllers

import (
	"log"

	"github.com/gerokkos/clerk/api/models"
	"github.com/jinzhu/gorm"
)

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Clerks{}, &models.Clerks{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.Clerks{}, &models.Clerks{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

}
