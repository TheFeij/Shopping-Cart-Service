package responses

import (
	"Shoping-Cart-Service/db/models"
	"time"
)

type GetBasketsListResponse struct {
	Baskets []GetBasketResponse `json:"baskets"`
}

type GetBasketResponse struct {
	ID        uint         `json:"id"`
	CreatedAt time.Time    `json:"createdAt"`
	UpdatedAt time.Time    `json:"updatedAt"`
	Data      []byte       `json:"data"`
	State     models.State `json:"state"`
}
