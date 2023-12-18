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
	ID        uint      `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
	Data      []byte    `gorm:"type:jsonb"`
	State     State     `gorm:"not null"`
}
