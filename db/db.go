package db

import (
	"Shoping-Cart-Service/db/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Init() {
	dsn := "host=localhost port=5432 user=root password=1234 dbname=shoping_cart_db sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
		return
	}

	if err := db.AutoMigrate(&models.ShoppingCart{}); err != nil {
		log.Fatalln(err)
		return
	}

	DB = db
}

func GetDB() *gorm.DB {
	return DB
}
