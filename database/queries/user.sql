-- name: GetUserById :one
SELECT * FROM users
WHERE id = $1;

-- name: GetUserByUserName :one
SELECT * FROM users
WHERE name = $1;

-- name: CreateUser :one
INSERT INTO users (name, pwd)
VALUES ($1, $2)
RETURNING id;