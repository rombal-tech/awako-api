package service

import (
	"alvile-api/models"
	"alvile-api/pkg/repository"
)

type SchemeService struct {
	repo repository.Scheme
}

func NewSchemeService(repo repository.Scheme) *SchemeService {
	return &SchemeService{repo: repo}
}

func (s *SchemeService) CreateScheme(schema models.Scheme, email string) (*models.Scheme, error) {
	return s.repo.CreateScheme(schema, email)
}

func (s *SchemeService) GetScheme(email string) ([]models.SchemeOutput, error) {
	return s.repo.GetScheme(email)
}
