package database

import (
	"api-rest/models"
	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DbConnection() {
	var err error
	DB, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Error to connection database: " + err.Error())
	} else {
		log.Println("DB Connected!")
	}
}

func RunDatabase() {
	//Connection
	DbConnection()
	//Migrations
	DB.AutoMigrate(models.Task{})
	DB.AutoMigrate(models.User{})
	DB.AutoMigrate(models.ActivityLog{})
}
