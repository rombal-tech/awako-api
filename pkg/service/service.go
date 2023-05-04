package service

import (
	"alvile-api/models"
	"alvile-api/pkg/repository"
)

type Registration interface {
	CreateAccount(email string, password string) (*models.Account, error)
	CreateSession(email string) (*models.Session, error)
}

type Account interface {
	IsExistByEmail(email string) (bool, error)
	CheckAuthorization(hed string) (string, error)
}

type Scheme interface {
	CreateScheme(scheme models.Scheme, email string) (*models.Scheme, error)
	GetScheme(email string) ([]models.SchemeOutput, error)
}

type Service struct {
	Registration
	Account
	Scheme
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Registration: NewAuthService(repos.Registration),
		Account:      NewAccountService(repos.Account),
		Scheme:       NewSchemeService(repos.Scheme),
	}
}
