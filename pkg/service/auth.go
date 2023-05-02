package service

import (
	"alvile-api/models"
	"alvile-api/pkg/repository"
	"alvile-api/pkg/service/generate"
)

type AuthService struct {
	repo repository.Registration
}

func NewAuthService(repo repository.Registration) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateAccount(email string, inputPassword string) (*models.Account, error) {
	var hash string
	var accountPassword string

	if inputPassword == "" {
		accountPassword = generate.Password()
	} else {
		accountPassword = inputPassword
	}

	hash = generate.PasswordHash(accountPassword)

	return s.repo.CreateAccount(email, hash)
}

func (s *AuthService) CreateSession(email string) (*models.Session, error) {
	hash := generate.Session()

	return s.repo.CreateSession(email, hash)
}
