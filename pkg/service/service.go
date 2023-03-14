package service

import (
	"alvile-api/models"
	"alvile-api/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.Account) (string, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
