package handlers

import (
	"Shoping-Cart-Service/db/models"
	"Shoping-Cart-Service/requests"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type BasketsHandler struct {
	db *gorm.DB
}

func NewBasketsHandler(db *gorm.DB) *BasketsHandler {
	return &BasketsHandler{
		db: db,
	}
}

func (handler BasketsHandler) CreateNewBasket(context *gin.Context) {

	var req requests.CreateNewBasketRequest
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	var newBasket models.Basket
	newBasket = models.Basket{
		Data:  req.Data,
		State: models.State(req.State),
	}

	if err := handler.db.Create(&newBasket).Error; err != nil {
		context.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	context.JSON(http.StatusOK, newBasket)
}

func errResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
