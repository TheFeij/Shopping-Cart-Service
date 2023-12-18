package requests

import "encoding/json"

type CreateNewBasketRequest struct {
	Data  json.RawMessage `json:"data" validate:"required,max=2048"`
	State string          `json:"state" validate:"required,in:PENDING,COMPLETED"`
}

type PatchBasketRequest struct {
	CreateNewBasketRequest
}
