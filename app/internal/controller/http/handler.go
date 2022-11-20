package http

import (
	"github.com/gin-gonic/gin"
)

type handler struct {
	// userBalance usecase.UserBalanse
}

func New() *handler {
	return &handler{}
	// return &handler{
	// 	userBalance: userBalance,
	// }
}

func (h *handler) NewRouter() *gin.Engine {
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
