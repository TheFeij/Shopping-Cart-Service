package main

import (
	"Shoping-Cart-Service/db"
	"database/sql"
	"log"
)

func main() {
	db.Init()
	db := db.GetDB()
	DB, err := db.DB()
	if err != nil {
		log.Fatalln(err)
	}
	defer func(DB *sql.DB) {
		err := DB.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(DB)

	//// sample data
	//basket := models.Basket{
	//	Data:  []byte(`{"name": "John Doe", "email": "johndoe@example.com"}`),
	//	State: models.COMPLETED,
	//}
	//db.Create(&basket)

	////retrieving sample data
	//var basket models.Basket
	//db.First(&basket, 1)
	//fmt.Println(string(basket.Data))
}
