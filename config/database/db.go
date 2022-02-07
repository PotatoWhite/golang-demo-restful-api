package database

import (
	"demo-restful-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var DB *gorm.DB

func Setup() {

	dsn := "host=localhost user=potato password=1234 dbname=test port=5432 sslmode=disable TimeZone=Asia/Seoul"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate([]models.Book{})
	DB = db
}

func GetDB() *gorm.DB {
	return DB
}
