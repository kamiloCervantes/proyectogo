package domain

import "github.com/google/uuid"

type URLRepository interface {
	Create(url *URL) error
	GetByID(id uuid.UUID) (*URL, error)
	GetAll() ([]URL, error)
	Update(url *URL) error
	Delete(id uuid.UUID) error
}
