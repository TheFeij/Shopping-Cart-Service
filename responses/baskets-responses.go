package responses

import (
	"Shoping-Cart-Service/db/models"
)

type GetBasketsListResponse struct {
	Baskets []models.Basket `json:"baskets"`
}

type GetBasketResponse struct {
	models.Basket
}
