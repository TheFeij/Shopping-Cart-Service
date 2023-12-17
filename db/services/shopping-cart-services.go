package services

import "gorm.io/gorm"

type ShoppingCartServices struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *ShoppingCartServices {
	return &ShoppingCartServices{
		DB: db,
	}
}
