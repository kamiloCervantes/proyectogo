package domain

import (
	"github.com/google/uuid"
)

type URL struct {
	ID            uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	Identificador string    `json:"identificador"`
	URL           string    `json:"url"`
	Etiquetas     string    `json:"etiquetas"`
}
