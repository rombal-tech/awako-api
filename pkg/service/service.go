package service

import (
	"alvile-api/models"
	"alvile-api/pkg/repository"
)

type Registration interface {
	CreateAccount(email string, password string) (*models.Account, error)
	CreateSession(email string) (*models.Session, error)
	CreateScheme(scheme models.Scheme, email string) (*models.SchemeOutput, error)
	CheckAuthorization(hed string) (string, error)
	GetScheme(email string) ([]models.SchemeOutput, error)
}

type Account interface {
	IsExistByEmail(email string) (bool, error)
}

type Service struct {
	Registration
	Account
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Registration: NewAuthService(repos.Registration),
		Account:      NewAccountService(repos.Account),
	}
}
