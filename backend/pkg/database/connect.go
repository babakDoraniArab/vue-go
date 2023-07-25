package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUsername,
		dbPassword,
		dbHost,
		dbPort,
		dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {

		log.Fatal("could not connect to the database")
		log.Fatal(err)
	}
	//DB is defined as a global variable in the common.go file
	fmt.Println("connected to the database")
	DB = db
}
