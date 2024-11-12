package domain

import (
	"github.com/google/uuid"
)

type URL struct {
	ID            uuid.UUID `gorm:"column:id;type:uuid;primaryKey;default:public.uuid_generate_v4()" json:"id"`
	Identificador string    `gorm:"column:identificador;type:varchar(500)" json:"identificador"`
	Url           string    `gorm:"column:dirurl;type:varchar(500)" json:"dirurl"`
	Etiquetas     string    `gorm:"column:etiquetas;type:varchar(500)" json:"etiquetas"`
}

func (URL) TableName() string {
	return "url"
}
