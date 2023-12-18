package models

import (
	"time"
)

type State string

const (
	COMPLETED State = "COMPLETED"
	PENDING   State = "PENDING"
)

type Basket struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `gorm:"not null" json:"createdAt"`
	UpdatedAt time.Time `gorm:"not null" json:"updatedAt"`
	Data      []byte    `gorm:"type:jsonb" json:"data"`
	State     State     `gorm:"not null" json:"state"`
}
