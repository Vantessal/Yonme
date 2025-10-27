package database

import (
	"accounts-service/internal/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dsn := "host=db port=5432 user=yonme_admin password=P@ssw0rd dbname=yonme_db"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	if err := db.AutoMigrate(&models.Account{}); err != nil {
		log.Fatalf("migration failed: %v", err)
	}
	return db
}