package service

import (
	"alvile-api/models"
	"alvile-api/pkg/repository"
)

type Registration interface {
	CreateUser(user models.Account) (string, error)
	CreateSession(session models.Session, email, password string) (string, error)
	CreateScheme(scheme models.Scheme) (int64, error)
	AuthorizationСheck(hed string) bool
}

type Service struct {
	Registration
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Registration: NewAuthService(repos.Registration),
	}
}
