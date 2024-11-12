package db

import (
	"proyectogo/internal/domain"
	"proyectogo/internal/infrastructure/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(cfg *config.Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(cfg.DatabaseDSN), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	// Migrar el esquema
	db.AutoMigrate(&domain.URL{})

	return db
}
