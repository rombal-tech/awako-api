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
