package database

import (
	"fmt"
	"log"
	"wira-backend/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=yourpassword dbname=wira_nft_store port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	fmt.Println("✅ Connected to PostgreSQL")

	DB.AutoMigrate(&models.NFT{}, &models.User{}, &models.Transaction{})
}
