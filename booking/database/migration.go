package database

import (
	"log"
	"microservices/pkg/entities"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Migrate(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&entities.Booking{})
	if err != nil {
		return nil, err
	}

	log.Println("✅ MySQL migration successful")
	return db, nil
}
