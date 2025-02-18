package database

import (
	"log"
	"fargo-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "null"

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	log.Println("Connected to the database")

	err = DB.AutoMigrate(&models.Client{}, &models.Admin{}, &models.CompanyContact{}, &models.TrackCode{})
	if err != nil {
		log.Fatal("Error during migration:", err)
	}

	log.Println("Database migrated successfully")
}
