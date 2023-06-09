package database

import (
	"OctavianoRyan25/GOSwagger/model"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "mulianet010"
	dbname   = "Car"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host = %s user = %v password = %v dbname = %v port = %v sslmode = disable", host, user, password, dbname, port)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database :", err)
	}

	db.Debug().AutoMigrate(model.Car{})
}

func GetDB() *gorm.DB {
	return db
}
