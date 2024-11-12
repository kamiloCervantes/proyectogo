package http

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, urlHandler *handler.URLHandler) {
	api := router.Group("/api")
	{
		api.POST("/urls", urlHandler.CreateURL)
		api.GET("/urls/:id", urlHandler.GetURLByID)
		api.GET("/urls", urlHandler.GetAllURLs)
		api.PUT("/urls", urlHandler.UpdateURL)
		api.DELETE("/urls/:id", urlHandler.DeleteURL)
	}
}
