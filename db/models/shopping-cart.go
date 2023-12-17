package models

import (
	"time"
)

type State string

const (
	COMPLETED State = "COMPLETED"
	PENDING   State = "PENDING"
)

type ShoppingCart struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Data      string `gorm:"type:varchar;size:2048;not null"`
	State     State  `gorm:"type:shopping_cart_state;not null"`
}
