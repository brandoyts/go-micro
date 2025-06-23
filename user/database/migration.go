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

	table := &entities.User{}

	// drop table if exists
	if db.Migrator().HasTable(table) {
		db.Migrator().DropTable(table)
	}

	err = db.AutoMigrate(table)
	if err != nil {
		return nil, err
	}

	log.Println("✅ MySQL migration successful")
	return db, nil
}

func Seed(db *gorm.DB) error {
	seedUser := entities.User{
		Username: "admin222",
		Password: "password", // Hash in production!
	}
	if err := db.Create(&seedUser).Error; err != nil {
		return err
	}
	log.Println("Inserted default user: admin")
	log.Println("🌱 Seeding complete")
	return nil
}
