// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: user_queries.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO Users (user_id, token, username, password, email, created_at, last_login)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING user_id, token, username, email, password, created_at, last_login
`

type CreateUserParams struct {
	UserID    uuid.UUID
	Token     uuid.UUID
	Username  string
	Password  string
	Email     string
	CreatedAt time.Time
	LastLogin time.Time
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.UserID,
		arg.Token,
		arg.Username,
		arg.Password,
		arg.Email,
		arg.CreatedAt,
		arg.LastLogin,
	)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Token,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.LastLogin,
	)
	return i, err
}

const deleteUserByToken = `-- name: DeleteUserByToken :exec
DELETE FROM Users WHERE token = $1
`

func (q *Queries) DeleteUserByToken(ctx context.Context, token uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteUserByToken, token)
	return err
}

const deleteUserByUsername = `-- name: DeleteUserByUsername :exec
DELETE FROM Users WHERE username = $1
`

func (q *Queries) DeleteUserByUsername(ctx context.Context, username string) error {
	_, err := q.db.Exec(ctx, deleteUserByUsername, username)
	return err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT user_id, token, username, email, password, created_at, last_login FROM Users WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Token,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.LastLogin,
	)
	return i, err
}

const getUserByToken = `-- name: GetUserByToken :one
SELECT user_id, token, username, email, password, created_at, last_login FROM Users WHERE token = $1
`

func (q *Queries) GetUserByToken(ctx context.Context, token uuid.UUID) (User, error) {
	row := q.db.QueryRow(ctx, getUserByToken, token)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Token,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.LastLogin,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT user_id, token, username, email, password, created_at, last_login FROM Users WHERE username = $1
`

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByUsername, username)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Token,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.LastLogin,
	)
	return i, err
}

const updateLastLoginTime = `-- name: UpdateLastLoginTime :exec
UPDATE Users SET last_login = $1 WHERE user_id = $2
`

type UpdateLastLoginTimeParams struct {
	LastLogin time.Time
	UserID    uuid.UUID
}

func (q *Queries) UpdateLastLoginTime(ctx context.Context, arg UpdateLastLoginTimeParams) error {
	_, err := q.db.Exec(ctx, updateLastLoginTime, arg.LastLogin, arg.UserID)
	return err
}
