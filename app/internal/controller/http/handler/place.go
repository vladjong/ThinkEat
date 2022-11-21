package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/vladjong/ThinkEat/internal/entities"
)

func (h *Handler) AddPlace(c *gin.Context) {
	var inputPlace entities.PlacePost
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
	places, err := h.placeUseCase.GetAllPlaces()
	if err != nil {
		logrus.Info(err)
		return
	}
	c.JSON(http.StatusOK, places)
}

func (h *Handler) GetPlace(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.Info(err)
		return
	}
	place, err := h.placeUseCase.GetPlace(id)
	if err != nil {
		logrus.Info(err)
		return
	}
	c.JSON(http.StatusOK, place)
}

func (h *Handler) UpdatePlace(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.Info(err)
		return
	}
	var inputPlace entities.PlacePost
	if err := c.BindJSON(&inputPlace); err != nil {
		logrus.Info(err)
		return
	}
	if err := h.placeUseCase.UpdatePlace(&inputPlace, id); err != nil {
		logrus.Info(err)
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "Ok",
	})
}

func (h *Handler) DeletePlace(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.Info(err)
		return
	}
	if err := h.placeUseCase.DeletePlace(id); err != nil {
		logrus.Info(err)
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "Ok",
	})
}
