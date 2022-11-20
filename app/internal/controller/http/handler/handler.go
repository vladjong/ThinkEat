package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/vladjong/ThinkEat/internal/domain"
)

type Handler struct {
	itemUseCase domain.ItemDomain
}

func New(item domain.ItemDomain) *Handler {
	return &Handler{
		itemUseCase: item,
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
			item.PUT("/:id", h.UpdateItem)
			item.DELETE("/:id", h.DeleteItem)
		}
		// place := api.Group("/place")
		// {
		// 	place.POST("/")
		// 	place.GET("/")
		// 	place.GET("/:id")
		// 	place.PUT("/:id")
		// 	place.DELETE("/:id")
		// }
	}
	return router
}
