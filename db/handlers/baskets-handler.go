package handlers

import (
	"Shoping-Cart-Service/db/models"
	"Shoping-Cart-Service/requests"
	"Shoping-Cart-Service/responses"
	"errors"
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

func (handler BasketsHandler) GetBasketsList(context *gin.Context) {
	var baskets []models.Basket

	if err := handler.db.Find(&baskets).Error; err != nil {
		context.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	context.JSON(http.StatusOK, responses.GetBasketsListResponse{Baskets: baskets})
}

func (handler BasketsHandler) GetBasket(context *gin.Context) {
	var basket models.Basket
	id := context.Param("id")

	if err := handler.db.First(&basket, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.JSON(http.StatusNotFound, gin.H{"error": "Basket not found"})
			return
		}

		context.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	context.JSON(http.StatusOK, responses.GetBasketResponse{Basket: basket})
}

func (handler BasketsHandler) DeleteBasket(context *gin.Context) {
	var basket models.Basket
	id := context.Param("id")

	if err := handler.db.Delete(&basket, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.JSON(http.StatusNotFound, gin.H{"error": "Basket not found"})
			return
		}

		context.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	context.JSON(http.StatusOK, basket)
}

func (handler BasketsHandler) UpdateBasket(context *gin.Context) {
	var basket models.Basket
	id := context.Param("id")

	if err := handler.db.First(&basket, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.JSON(http.StatusNotFound, gin.H{"error": "Basket not found"})
			return
		}

		context.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	var req requests.PatchBasketRequest
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	basket.Data = req.Data
	basket.State = models.State(req.State)

	if err := handler.db.Save(&basket).Error; err != nil {
		context.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	context.JSON(http.StatusOK, basket)
}

func errResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
