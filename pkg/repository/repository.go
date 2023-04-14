package repository

import (
	"alvile-api/models"
	"github.com/jmoiron/sqlx"
)

type Registration interface {
	CreateUser(user models.AccountInput) (*models.AccountRegistrationOutput, error)
	GetUser(email, password string) (string, error)
	CreateSession(session *models.Session) (*models.SessionOutput, error)
	CreateScheme(schema models.Scheme, email string) (*models.SchemeOutput, error)
	CheckAuthorization(hed string) (string, error)
	GetScheme(email string) ([]models.SchemeOutput, error)
}

type Repository struct {
	Registration
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Registration: NewAuthPostgres(db),
	}
}
