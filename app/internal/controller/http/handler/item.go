package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vladjong/ThinkEat/internal/controller/http/handler/dto"
)

func (h *Handler) AddItem(c *gin.Context) {
	var inputItem dto.ItemDto
	if err := c.BindJSON(&inputItem); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.itemUseCase.AddItem(inputItem); err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "Ok",
	})
}

func (h *Handler) GetAllItems(c *gin.Context) {
	items, err := h.itemUseCase.GetAllItems()
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, items)

}

func (h *Handler) GetItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	item, err := h.itemUseCase.GetItem(id)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) UpdateItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	var inputItem dto.ItemDto
	if err := c.BindJSON(&inputItem); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.itemUseCase.UpdateItem(inputItem, id); err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "Ok",
	})
}

func (h *Handler) DeleteItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.itemUseCase.DeleteItem(id); err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "Ok",
	})
}
