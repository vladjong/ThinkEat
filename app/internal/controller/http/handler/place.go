package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/vladjong/ThinkEat/internal/entities"
)

func (h *Handler) AddPlace(c *gin.Context) {
	var inputPlace entities.Place
	if err := c.BindJSON(&inputPlace); err != nil {
		logrus.Info(err)
		return
	}
	if err := h.placeUseCase.AddPlace(&inputPlace); err != nil {
		logrus.Info(err)
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "Ok",
	})
}

func (h *Handler) GetAllPlaces(c *gin.Context) {

}

func (h *Handler) GetPlace(c *gin.Context) {

}

func (h *Handler) UpdatePlace(c *gin.Context) {

}

func (h *Handler) DeletePlace(c *gin.Context) {

}
