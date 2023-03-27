package repository

import (
	"alvile-api/models"
	"github.com/jmoiron/sqlx"
)

type Registration interface {
	CreateUser(user models.Account) (string, error)
	GetUser(email, password string) (string, error)
	CreateSession(session models.Session) (string, error)
}

type Repository struct {
	Registration
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Registration: NewAuthPostgres(db),
	}
}
