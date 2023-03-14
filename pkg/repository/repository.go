package repository

import (
	"alvile-api/models"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user models.Account) (string, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
