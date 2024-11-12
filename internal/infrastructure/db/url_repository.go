package db

import (
	"proyectogo/internal/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type urlRepository struct {
	db *gorm.DB
}

func NewURLRepository(db *gorm.DB) domain.URLRepository {
	return &urlRepository{db}
}

func (r *urlRepository) Create(url *domain.URL) error {
	return r.db.Create(url).Error
}

func (r *urlRepository) GetByID(id uuid.UUID) (*domain.URL, error) {
	var url domain.URL
	if err := r.db.First(&url, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &url, nil
}

func (r *urlRepository) GetAll() ([]domain.URL, error) {
	var urls []domain.URL
	if err := r.db.Find(&urls).Error; err != nil {
		return nil, err
	}
	return urls, nil
}

func (r *urlRepository) Update(url *domain.URL) error {
	return r.db.Save(url).Error
}

func (r *urlRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&domain.URL{}, "id = ?", id).Error
}
