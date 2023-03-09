package repository

import (
	"alvile-api/models"
	"alvile-api/queries"
)

type Account interface {
	IsExistByEmail(email string) (bool, error)
	IsExistByID(id int64) (bool, error)
	Registration(input *models.RegistrationAccountInput) (*queries.Account, error)
}

type Repository struct {
	Account
}

func NewRepository(db *queries.Queries) *Repository {
	return &Repository{
		Account: NewAccountPostgres(db),
	}
}
