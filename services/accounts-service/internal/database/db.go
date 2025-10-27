package database

import (
	"accounts-service/internal/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dsn := os.Getenv("DATABASE_URL") // например: postgres://user:pass@localhost:5432/accounts
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	if err := db.AutoMigrate(&models.Account{}); err != nil {
		log.Fatalf("migration failed: %v", err)
	}
	return db
}