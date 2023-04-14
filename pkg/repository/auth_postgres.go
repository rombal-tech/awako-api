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

func (r *AuthPostgres) CreateUser(user models.AccountInput) (*models.AccountRegistrationOutput, error) {
	var output models.AccountRegistrationOutput
	query := fmt.Sprintf("INSERT INTO \"public.Account\" (email,password,deleted) values($1,$2,false)RETURNING email")
	row := r.db.QueryRow(query, user.Email, user.Password)
	if err := row.Scan(&output.Email); err != nil {
		return nil, err
	}

	return &output, nil
}

func (r *AuthPostgres) GetUser(email, password string) (string, error) {
	var userEmail string
	query := fmt.Sprintf("SELECT email FROM \"public.Account\" WHERE email =$1 AND password =$2")
	err := r.db.Get(&userEmail, query, email, password)
	return userEmail, err
}

func (r *AuthPostgres) CreateSession(session *models.Session) (*models.SessionOutput, error) {
	var sessionString models.SessionOutput
	query := fmt.Sprintf("INSERT INTO \"public.Session\" (email,session_string) values($1,$2)RETURNING session_string")
	row := r.db.QueryRow(query, session.Email, session.SessionString)
	if err := row.Scan(&sessionString.SessionString); err != nil {
		return nil, err
	}

	return &sessionString, nil
}

func (r *AuthPostgres) CreateScheme(schema models.Scheme, email string) (*models.SchemeOutput, error) {
	var scheme models.SchemeOutput
	query := fmt.Sprintf("INSERT INTO \"public.Scheme\" (name,description,author,creation_date) values($1,$2,$3,CURRENT_TIMESTAMP) RETURNING id")
	row := r.db.QueryRow(query, schema.Name, schema.Description, email)
	if err := row.Scan(&scheme.Id); err != nil {
		return &scheme, err
	}

	return &scheme, nil
}

func (r *AuthPostgres) CheckAuthorization(hed string) (string, error) {
	var email string
	query := fmt.Sprintf("SELECT email FROM \"public.Session\" WHERE session_string =$1 ")
	err := r.db.Get(&email, query, hed)
	return email, err
}

func (r *AuthPostgres) GetScheme(email string) ([]models.SchemeOutput, error) {
	var output []models.SchemeOutput
	query := fmt.Sprintf("SELECT id,name,description,author,creation_date FROM \"public.Scheme\" WHERE author =$1 ")
	err := r.db.Select(&output, query, email)
	return output, err
}
