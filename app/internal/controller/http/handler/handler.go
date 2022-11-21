package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/vladjong/ThinkEat/internal/domain"
)

type Handler struct {
	itemUseCase  domain.ItemDomain
	placeUseCase domain.PlaceDomain
}

func New(item domain.ItemDomain, place domain.PlaceDomain) *Handler {
	return &Handler{
		itemUseCase:  item,
		placeUseCase: place,
	}
}

func (h *Handler) NewRouter() *gin.Engine {
	router := gin.New()

	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")
	{
		item := api.Group("/item")
		{
			item.POST("/", h.AddItem)
			item.GET("/", h.GetAllItems)
			item.GET("/:id", h.GetItem)
			item.DELETE("/:id", h.DeleteItem)
		}
		place := api.Group("/place")
		{
			place.POST("/", h.AddPlace)
			place.GET("/", h.GetAllPlaces)
			place.GET("/:id", h.GetPlace)
			place.PUT("/:id", h.UpdatePlace)
			place.DELETE("/:id", h.DeletePlace)
		}
	}
	return router
}
