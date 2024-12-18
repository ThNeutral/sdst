// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package plasmadb

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	AddUserToProject(ctx context.Context, arg AddUserToProjectParams) error
	CreateMessage(ctx context.Context, arg CreateMessageParams) (Message, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteUserByToken(ctx context.Context, token uuid.UUID) error
	DeleteUserByUsername(ctx context.Context, username string) error
	DeleteUserFromProject(ctx context.Context, arg DeleteUserFromProjectParams) error
	GetMessagesByProject(ctx context.Context, projectID uuid.UUID) ([]Message, error)
	GetMessagesByUsers(ctx context.Context, userID uuid.UUID) ([]Message, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUserById(ctx context.Context, arg GetUserByIdParams) (ProjectUser, error)
	GetUserByToken(ctx context.Context, token uuid.UUID) (User, error)
	GetUserByUsername(ctx context.Context, username string) (User, error)
	UpdateLastLoginTime(ctx context.Context, arg UpdateLastLoginTimeParams) error
	UpdateUserRole(ctx context.Context, arg UpdateUserRoleParams) error
}

var _ Querier = (*Queries)(nil)
