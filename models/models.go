package models

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var err error

func Connect(connectionString string) {
	Instance, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database...")
}
func Migrate() {
	Instance.AutoMigrate(&Student{})
	log.Println("Database Migration Completed...")
}

type Student struct {
	UseID    int    `gorm:"primaryKey"`
	Username string `json:"user_name"`
	Address  string `json:"user_address"`
}
