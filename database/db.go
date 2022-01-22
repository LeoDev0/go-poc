package database

import (
	"fmt"
	"log"

	"example.com/go-poc/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	err error
)

func DatabaseConnection(connectionString string) {
	DB, err = gorm.Open(postgres.Open(connectionString))
	if err != nil {
		log.Panic("Error to connect to connect to database: ", err)
	} else {
		fmt.Println("Succesfully connected to database")
	}

	DB.AutoMigrate(&models.Book{})
	DB.AutoMigrate(&models.Author{})
}