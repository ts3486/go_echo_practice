package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB
var e error

func DatabaseInit() {
	host :=  "localhost"
	user := "postgres"
	password := "86Zms98"
	dbName := "test"
	port := 5432

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbName, port)
	database, e = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if e != nil {
		fmt.Println("Connection NOT Opened to Database")
		panic(e)
	}
}

func Connect() *gorm.DB {
	DatabaseInit()
	fmt.Println("Connection Opened to Database")
	return database
}