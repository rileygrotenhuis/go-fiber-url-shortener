package database

import (
	"go-fiber-url-shortener/model"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() {
	var err error

	env := os.Getenv("DATABASE_URL")

	DB, err = gorm.Open(postgres.Open(env), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	err = DB.AutoMigrate(&model.Goly{})

	if err != nil {
		log.Fatal(err)
	}
}
