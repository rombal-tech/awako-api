package repository

import (
	"alvile-api/models"
	"github.com/jmoiron/sqlx"
)

type Registration interface {
	CreateAccount(email, password string) (*models.Account, error)
	CreateSession(email, hash string) (*models.Session, error)
	CreateScheme(schema models.Scheme, email string) (*models.SchemeOutput, error)
	CheckAuthorization(hed string) (string, error)
	GetScheme(email string) ([]models.SchemeOutput, error)
}

type Account interface {
	IsExistByEmail(email string) (bool, error)
}

type Repository struct {
	Registration
	Account
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Registration: NewAuthPostgres(db),
		Account:      NewAccountPostgres(db),
	}
}
