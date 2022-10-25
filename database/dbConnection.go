package database

import (
	"CRUDRestAPI/model"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var connector *gorm.DB

var dbURL = "root:BbSs@563884(localhost:3306)/contactdb"

func init() {
	var err error

	connector, err = gorm.Open(mysql.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatal("Error in DB Connection", err)
	}
	log.Printf("Connection Successful...")
	connector.AutoMigrate(&model.Contact{})
}
