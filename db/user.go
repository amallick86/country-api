package db

import (
	"context"
	"github.com/amallick86/country-api/models"
)

// create user in database
const createUser = `
INSERT INTO users (
  username, password
) VALUES (
  $1, $2
)
RETURNING id, username, password, created_at
`

type CreateUserParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (models.User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Username, arg.Password)
	var i models.User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.CreatedAt,
	)
	return i, err
}

// Get user for login

const getUser = `
SELECT id, username, password, created_at FROM users
WHERE username = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, username string) (models.User, error) {
	row := q.db.QueryRowContext(ctx, getUser, username)
	var i models.User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.CreatedAt,
	)
	return i, err
}

// Get user by id

const getUserById = `
SELECT id, username, password, created_at FROM users
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUserById(ctx context.Context, userId int32) (models.User, error) {
	row := q.db.QueryRowContext(ctx, getUserById, userId)
	var i models.User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.CreatedAt,
	)
	return i, err
}
