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
	dsn := os.Getenv("USER_DB_DSN")
	log.Println("Connecting to MySQL with DSN:", dsn)
	if dsn == "" {
		log.Fatal("USER_DB_DSN environment variable not set")
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return &Database{Db: db}

}

func (d *Database) Create(user entities.User) *entities.User {
	result := d.Db.Create(&user)

	if result.Error != nil {
		log.Fatal(result.Error)
		return nil
	}

	return &user
}
