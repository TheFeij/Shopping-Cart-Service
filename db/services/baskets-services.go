package services

import "gorm.io/gorm"

type BasketsServices struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *BasketsServices {
	return &BasketsServices{
		DB: db,
	}
}
