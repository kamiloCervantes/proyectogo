package app

import (
	"proyectogo/internal/domain"

	"github.com/google/uuid"
)

type URLService interface {
	CreateURL(url *domain.URL) error
	GetURLByID(id uuid.UUID) (*domain.URL, error)
	GetAllURLs() ([]domain.URL, error)
	UpdateURL(url *domain.URL) error
	DeleteURL(id uuid.UUID) error
}

type urlService struct {
	repo domain.URLRepository
}

func NewURLService(repo domain.URLRepository) URLService {
	return &urlService{repo}
}

func (s *urlService) CreateURL(url *domain.URL) error {
	url.ID = uuid.New()
	return s.repo.Create(url)
}

func (s *urlService) GetURLByID(id uuid.UUID) (*domain.URL, error) {
	return s.repo.GetByID(id)
}

func (s *urlService) GetAllURLs() ([]domain.URL, error) {
	return s.repo.GetAll()
}

func (s *urlService) UpdateURL(url *domain.URL) error {
	return s.repo.Update(url)
}

func (s *urlService) DeleteURL(id uuid.UUID) error {
	return s.repo.Delete(id)
}
