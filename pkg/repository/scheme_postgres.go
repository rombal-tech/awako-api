package repository

import (
	"alvile-api/errors"
	"alvile-api/models"
	"github.com/execaus/exloggo"
	"github.com/jmoiron/sqlx"
)

type SchemePostgres struct {
	db *sqlx.DB
}

func NewSchemePostgres(db *sqlx.DB) *SchemePostgres {
	return &SchemePostgres{db: db}
}

func (r *SchemePostgres) CreateScheme(inputSchema models.Scheme, email string) (*models.Scheme, error) {
	var scheme models.Scheme
	row := r.db.QueryRow("INSERT INTO \"Scheme\" (name,description,author,creation_date) values($1,$2,$3,CURRENT_TIMESTAMP) RETURNING *",
		inputSchema.Name, inputSchema.Description, email)
	if err := row.Scan(&scheme.Id, &scheme.Name, &scheme.Description, &scheme.Author, &scheme.CreationDate); err != nil {
		exloggo.Error(err.Error())
		return nil, errors.ServerError
	}

	return &scheme, nil
}

func (r *SchemePostgres) GetScheme(email string) ([]models.SchemeOutput, error) {
	var output []models.SchemeOutput
	err := r.db.Select(output, `SELECT * FROM "Scheme" WHERE author =$1 `, email)
	if err != nil {
		exloggo.Error(err.Error())
		return nil, errors.ServerError
	}
	return output, err
}
