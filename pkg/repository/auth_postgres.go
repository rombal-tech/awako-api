package repository

import (
	"alvile-api/errors"
	"alvile-api/models"
	"github.com/execaus/exloggo"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateAccount(email string, password string) (*models.Account, error) {
	var account models.Account

	if err := r.db.Get(&account, "INSERT INTO \"Account\" (email,password,deleted,confirmed,confirm_code) VALUES ($1,$2,false,true,1234) RETURNING *",
		email,
		password); err != nil {
		exloggo.Error(err.Error())
		return nil, errors.ServerError
	}

	return &account, nil
}

func (r *AuthPostgres) CreateSession(email, hash string) (*models.Session, error) {
	var session models.Session

	if err := r.db.Get(&session,
		`INSERT INTO "Session" (email, session) VALUES ($1, $2) RETURNING *`, email, hash); err != nil {
		exloggo.Error(err.Error())
		return nil, errors.ServerError
	}

	return &session, nil
}
