package repository

import (
	"alvile-api/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user models.Account) (string, error) {
	var email string
	query := fmt.Sprintf("INSERT INTO \"public.Account\" (email,password,deleted) values($1,$2,false)RETURNING email")
	row := r.db.QueryRow(query, user.Email, user.Password)
	if err := row.Scan(&email); err != nil {
		return "", err
	}

	return email, nil
}

func (r *AuthPostgres) GetUser(email, password string) (string, error) {
	var userEmail string
	query := fmt.Sprintf("SELECT email FROM \"public.Account\" WHERE email =$1 AND password =$2")
	err := r.db.Get(&userEmail, query, email, password)
	return userEmail, err
}

func (r *AuthPostgres) CreateSession(session models.Session) (string, error) {
	var sessionString string
	query := fmt.Sprintf("INSERT INTO \"public.Session\" (email,session_string) values($1,$2)RETURNING session_string")
	row := r.db.QueryRow(query, session.Email, session.SessionString)
	if err := row.Scan(&sessionString); err != nil {
		return "", err
	}

	return sessionString, nil
}
