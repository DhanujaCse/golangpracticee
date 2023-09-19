package controllers

import (
	"jwtEx/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var dbError error

func Connect() {
	database := ("root:Dhanu123@tcp(localhost:3306)/prg?parseTime=true")
	Instance, dbError = gorm.Open(mysql.Open(database), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to DB")
	}
	Instance.AutoMigrate(&models.Userlogindetails{})
	log.Println("Connected to Database!")
}
