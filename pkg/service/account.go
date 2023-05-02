package service

import "alvile-api/pkg/repository"

type AccountService struct {
	repo repository.Account
}

func NewAccountService(repo repository.Account) *AccountService {
	return &AccountService{repo: repo}
}

func (s *AccountService) IsExistByEmail(email string) (bool, error) {
	return s.repo.IsExistByEmail(email)
}

func (s *AccountService) CheckAuthorization(hed string) (string, error) {
	return s.repo.CheckAuthorization(hed)
}
