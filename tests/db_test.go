package tests

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/gerokkos/clerk/api/controllers"
	"github.com/gerokkos/clerk/api/models"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var server = controllers.Server{}
var userInstance = models.User{}

func TestMain(m *testing.M) {
	var err error
	err = godotenv.Load(os.ExpandEnv(".env"))
	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}
	Database()

	log.Printf("Before calling m.Run() !!!")
	ret := m.Run()
	log.Printf("After calling m.Run() !!!")
	//os.Exit(m.Run())
	os.Exit(ret)
}

func Database() {

	var err error

	TestDbDriver := os.Getenv("TestDbDriver")

	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("TestDbHost"), os.Getenv("TestDbPort"), os.Getenv("TestDbUser"), os.Getenv("TestDbName"), os.Getenv("TestDbPassword"))
	server.DB, err = gorm.Open(TestDbDriver, DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to %s database\n", TestDbDriver)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database\n", TestDbDriver)
	}
}

func refreshUserTable() error {
	server.DB.Exec("SET foreign_key_checks=0")
	err := server.DB.Debug().DropTableIfExists(&models.User{}).Error
	if err != nil {
		return err
	}
	server.DB.Exec("SET foreign_key_checks=1")
	err = server.DB.Debug().AutoMigrate(&models.User{}).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed table")
	log.Printf("refreshUserTable routine OK !!!")
	return nil
}
