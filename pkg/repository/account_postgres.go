package repository

import (
	"alvile-api/models"
	"alvile-api/queries"
	"github.com/sirupsen/logrus"
)

type AccountPostgres struct {
	db *queries.Queries
}

func (r *AccountPostgres) IsExistByEmail(email string) (bool, error) {
	isExist, err := r.db.IsExistAccountByEmail(context.Background(), email)
	if err != nil {
		logrus.Error(err.Error())
		return false, err
	}
	return isExist, nil
}

func (r *AccountPostgres) IsExistByID(id int64) (bool, error) {
	isExist, err := r.db.IsExistAccountByID(context.Background(), id)
	if err != nil {
		logrus.Error(err.Error())
		return false, err
	}
	return isExist, nil
}

func (r *AccountPostgres) Registration(input *models.RegistrationAccountInput) (*queries.Account, error) {
	account, err := r.db.CreateAccount(context.Background(), queries.CreateAccountParams{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Password:  input.Password,
	})
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return &account, nil
}

func NewAccountPostgres(db *queries.Queries) *AccountPostgres {
	return &AccountPostgres{db: db}
}
