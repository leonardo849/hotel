package repository

import (
	"fmt"
	"hotel/internal/model"
	"log"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() (*gorm.DB, error) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		return nil, fmt.Errorf("there isn't dsn")
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	DB = db
	err = migrateModels(db)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func migrateModels(db *gorm.DB) error {
	err := db.AutoMigrate(&model.Guest{}, &model.Room{}, &model.Reservation{})
	if err != nil {
		return err
	}
	log.Println("Tables are ok")
	return nil
}

