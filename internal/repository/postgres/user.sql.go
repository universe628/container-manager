// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: user.sql

package postgresRepo

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (name, pwd)
VALUES ($1, $2)
RETURNING id
`

type CreateUserParams struct {
	Name string
	Pwd  string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Name, arg.Pwd)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const getUserById = `-- name: GetUserById :one
SELECT id, name, pwd, created_at, updated_at, deleted_at FROM users
WHERE id = $1
`

func (q *Queries) GetUserById(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Pwd,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getUserByUserName = `-- name: GetUserByUserName :one
SELECT id, name, pwd, created_at, updated_at, deleted_at FROM users
WHERE name = $1
`

func (q *Queries) GetUserByUserName(ctx context.Context, name string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByUserName, name)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Pwd,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}
