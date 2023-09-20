-- name: CreateUser :one
INSERT INTO users (id, username, email, password) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetAllUsers :many
SELECT * FROM users;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;