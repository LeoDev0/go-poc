package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	err error
)

func DatabaseConnection(conn string) {
	connectionString := conn
	DB, err = gorm.Open(postgres.Open(connectionString))
	if err != nil {
		log.Panic("Error to connect to connect to database: ", err)
	} else {
		fmt.Println("Succesfully connected to database")
	}
}