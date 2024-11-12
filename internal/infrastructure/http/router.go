package http

import (
	"proyectogo/internal/app/handler"

	e "proyectogo/internal/utils/env"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, urlHandler *handler.URLHandler) {

	Env := e.NewEnv()
	Env.Env()
	api := router.Group("/api")
	{

		if Env.Mode == "QUERY" {
			api.POST("/urls", urlHandler.CreateURL)
			api.PUT("/urls", urlHandler.UpdateURL)
			api.DELETE("/urls/:id", urlHandler.DeleteURL)
		} else {
			api.GET("/urls/:id", urlHandler.GetURLByID)
			api.GET("/urls", urlHandler.GetAllURLs)
		}

	}
}
