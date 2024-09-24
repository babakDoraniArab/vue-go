package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {

	dbUsername := "example"
	dbPassword := "example"
	dbHost := "localhost"
	dbPort := "3306"
	dbName := "example"
	fmt.Println("this is the dbUsername:", dbUsername)
	fmt.Println("this is the dbPassword:", dbPassword)
	fmt.Println("this is the dbHost:", dbHost)
	fmt.Println("this is the dbPort:", dbPort)
	fmt.Println("this is the dbName:", dbName)
	fmt.Println("this is the dbPort:", dbPort)
	//dns="freedb_gobabak":"AyaRn8%ymVPH!uA"@tcp(sql.freedb.tech:3306)/freedb_gobabak?charset=utf8mb4&parseTime=True&loc=Local
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
