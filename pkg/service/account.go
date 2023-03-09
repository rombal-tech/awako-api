package service

import (
	"alvile-api/models"
	"alvile-api/pkg/repository"
)

const (
	accountGetListDefaultLimit  = 10
	accountGetListDefaultOffset = 0
)

type AccountService struct {
	repo repository.Account
}

func (s *AccountService) IsExistByEmail(email string) (bool, error) {
	return s.repo.IsExistByEmail(email)
}

func (s *AccountService) IsExistByID(id int64) (bool, error) {
	return s.repo.IsExistByID(id)
}

func (s *AccountService) Registration(
	input *models.RegistrationAccountInput,
) (*models.RegistrationAccountOutput, error) {
	account, err := s.repo.Registration(input)
	if err != nil {
		return nil, err
	}

	return &models.RegistrationAccountOutput{
		ID:        account.ID,
		FirstName: account.FirstName,
		LastName:  account.LastName,
		Email:     account.Email,
	}, nil
}

func NewAccountService(repo repository.Account) *AccountService {
	return &AccountService{repo: repo}
}
