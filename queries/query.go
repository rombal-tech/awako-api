package queries

import (
	"context"
)

const createAccount = `-- name: CreateAccount :one
INSERT INTO "Account" (first_name, last_name, email, password, deleted)
VALUES ($1, $2, $3, $4, false)
RETURNING id, first_name, last_name, email, password, deleted
`

type CreateAccountParams struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, createAccount,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.Password,
	)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Password,
		&i.Deleted,
	)
	return i, err
}
