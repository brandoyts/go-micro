package database

import (
	"log"
	"microservices/pkg/entities"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	Db *gorm.DB
}

func New() *Database {
	dsn := os.Getenv("BOOKING_DB_DSN")
	log.Println("Connecting to MySQL with DSN:", dsn)
	if dsn == "" {
		log.Fatal("BOOKING_DB_DSN environment variable not set")
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: initLogger(),
	})
	if err != nil {
		log.Fatal(err)
	}

	return &Database{Db: db}

}

func (d *Database) Create(booking entities.Booking) *entities.Booking {
	result := d.Db.Create(&booking)

	if result.Error != nil {
		log.Fatal(result.Error)
		return nil
	}

	return &booking
}
