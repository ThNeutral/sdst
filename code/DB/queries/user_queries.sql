-- name: CreateUser :one
INSERT INTO Users (user_id, first_name, last_name, password, email, created_at, last_login, role_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;