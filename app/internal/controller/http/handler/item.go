package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/vladjong/ThinkEat/internal/entities"
)

func (h *Handler) AddItem(c *gin.Context) {
	var inputItem entities.Item
	if err := c.BindJSON(&inputItem); err != nil {
		logrus.Info(err)
		return
	}
	if err := h.itemUseCase.AddItem(&inputItem); err != nil {
		logrus.Info(err)
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "Ok",
	})
}

func (h *Handler) GetAllItems(c *gin.Context) {

}

func (h *Handler) GetItem(c *gin.Context) {

}

func (h *Handler) UpdateItem(c *gin.Context) {

}

func (h *Handler) DeleteItem(c *gin.Context) {

}
