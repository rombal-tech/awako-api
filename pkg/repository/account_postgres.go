package repository

import (
	"alvile-api/errors"
	"github.com/execaus/exloggo"
	"github.com/jmoiron/sqlx"
)

type AccountPostgres struct {
	db *sqlx.DB
}

func NewAccountPostgres(db *sqlx.DB) *AccountPostgres {
	return &AccountPostgres{db: db}
}

func (r *AccountPostgres) IsExistByEmail(email string) (bool, error) {
	var isExist bool

	if err := r.db.Get(&isExist, `SELECT EXISTS (SELECT 1 FROM "Account" WHERE email=$1 AND deleted=false`,
		email); err != nil {
		exloggo.Error(err.Error())
		return false, errors.ServerError
	}

	return isExist, nil
}

func (r *AccountPostgres) CheckAuthorization(hed string) (string, error) {
	var email string
	err := r.db.Get(&email, `SELECT email FROM "Session" WHERE session_string =$1`, hed)
	if err != nil {
		exloggo.Error(err.Error())
		return "", errors.ServerError
	}
	return email, nil
}
