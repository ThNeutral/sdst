-- name: CreateUser :one
INSERT INTO Users (user_id, token, username, password, email, created_at, last_login)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: UpdateLastLoginTime :exec
UPDATE Users SET last_login = $1 WHERE user_id = $2;

-- name: GetUserByEmail :one
SELECT * FROM Users WHERE email = $1;

-- name: GetUserByUsername :one
SELECT * FROM Users WHERE username = $1;

-- name: GetUserByToken :one
SELECT * FROM Users WHERE token = $1;

-- name: DeleteUserByToken :exec
DELETE FROM Users WHERE token = $1;

-- name: DeleteUserByUsername :exec
DELETE FROM Users WHERE username = $1;