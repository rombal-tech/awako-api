package repository

import (
	"alvile-api/errors"
	"alvile-api/models"
	"fmt"
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

	if err := r.db.Get(`INSERT INTO "Account" (email,password,deleted) VALUES ($1,$2,false) RETURNING *`,
		email,
		password); err != nil {
		exloggo.Error(err.Error())
		return nil, errors.ServerError
	}

	return &account, nil
}

func (r *AuthPostgres) GetAccount(email, password string) (string, error) {
	var userEmail string
	query := fmt.Sprintf("SELECT email FROM \"Account\" WHERE email =$1 AND password =$2")
	err := r.db.Get(&userEmail, query, email, password)
	return userEmail, err
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

func (r *AuthPostgres) CreateScheme(schema models.Scheme, email string) (*models.SchemeOutput, error) {
	var scheme models.SchemeOutput
	query := fmt.Sprintf("INSERT INTO \"Scheme\" (name,description,author,creation_date) values($1,$2,$3,CURRENT_TIMESTAMP) RETURNING id")
	row := r.db.QueryRow(query, schema.Name, schema.Description, email)
	if err := row.Scan(&scheme.Id); err != nil {
		return &scheme, err
	}

	return &scheme, nil
}

func (r *AuthPostgres) CheckAuthorization(hed string) (string, error) {
	var email string
	query := fmt.Sprintf("SELECT email FROM \"Session\" WHERE session_string =$1 ")
	err := r.db.Get(&email, query, hed)
	return email, err
}

func (r *AuthPostgres) GetScheme(email string) ([]models.SchemeOutput, error) {
	var output []models.SchemeOutput
	query := fmt.Sprintf("SELECT id,name,description,author,creation_date FROM \"Scheme\" WHERE author =$1 ")
	err := r.db.Select(&output, query, email)
	return output, err
}
