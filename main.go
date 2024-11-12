package main

import (
	"proyectogo/internal/app"
	"proyectogo/internal/app/handler"
	"proyectogo/internal/infrastructure/config"
	"proyectogo/internal/infrastructure/db"
	"proyectogo/internal/infrastructure/http"
	"proyectogo/internal/infrastructure/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	log := logger.NewLogger()

	dbConn := db.NewDatabase(cfg)

	urlRepo := db.NewURLRepository(dbConn)

	urlService := app.NewURLService(urlRepo)

	urlHandler := handler.NewURLHandler(urlService)

	router := gin.Default()
	http.RegisterRoutes(router, urlHandler)

	log.Println("Iniciando server en " + cfg.ServerAddress)
	router.Run(cfg.ServerAddress)
}
