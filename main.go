package main

import (
	"proyectogo/internal/app"
	"proyectogo/internal/infrastructure/config"
	"proyectogo/internal/infrastructure/db"
	"proyectogo/internal/infrastructure/http"
	"proyectogo/internal/infrastructure/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	// Cargar configuración
	cfg := config.LoadConfig()

	// Inicializar logger
	log := logger.NewLogger()

	// Conectar a la base de datos
	dbConn := db.NewDatabase(cfg)

	// Inicializar repositorio
	urlRepo := db.NewURLRepository(dbConn)

	// Inicializar servicio
	urlService := app.NewURLService(urlRepo)

	// Inicializar manejador
	urlHandler := app.NewURLHandler(urlService)

	// Configurar enrutador
	router := gin.Default()
	http.RegisterRoutes(router, urlHandler)

	// Iniciar servidor
	log.Info("Starting server at " + cfg.ServerAddress)
	router.Run(cfg.ServerAddress)
}
