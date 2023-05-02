package repository

import (
	"alvile-api/errors"
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
		return false, errors.ServerError
	}

	return isExist, nil
}
