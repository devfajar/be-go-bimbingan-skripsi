package database

import (
	"github.com/devfajar/go-bimbingan-skripsi/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() (*gorm.DB, error) {
	// DB Config
	dsn := "host=localhost user=root password=bimbing4n dbname=bimbingan port=5432 sslmode=disable TimeZone=Asia/Jakarta"

	// Connect to DB
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	DB = db

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
