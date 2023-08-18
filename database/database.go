package database

import (
	"fmt"
	"golang_jwt/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbname   = "golang"
	host     = "localhost"
	port     = "5432"
	username = "postgres"
	password = "password"
	db       *gorm.DB
	err      error
)

func StartDB() {

	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s password=%s sslmode=disable", host, username, password, dbname, password)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		log.Fatal("error database :", err)
	}

	db.Debug().AutoMigrate(models.User{}, models.Product{})

}

func GetDB() *gorm.DB {
	return db
}
