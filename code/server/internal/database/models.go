// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"database/sql/driver"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type UserRoles string

const (
	UserRolesSuperAdmin UserRoles = "Super Admin"
	UserRolesAdmin      UserRoles = "Admin"
	UserRolesUser       UserRoles = "User"
)

func (e *UserRoles) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UserRoles(s)
	case string:
		*e = UserRoles(s)
	default:
		return fmt.Errorf("unsupported scan type for UserRoles: %T", src)
	}
	return nil
}

type NullUserRoles struct {
	UserRoles UserRoles
	Valid     bool // Valid is true if UserRoles is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullUserRoles) Scan(value interface{}) error {
	if value == nil {
		ns.UserRoles, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.UserRoles.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullUserRoles) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.UserRoles), nil
}

type Project struct {
	ProjectID   uuid.UUID
	PName       string
	Description string
	OwnerID     int32
	CreatedAt   pgtype.Timestamp
	UpdatedAt   pgtype.Timestamp
}

type ProjectFile struct {
	FileID       uuid.UUID
	FileLocation pgtype.Text
	FileContent  pgtype.Text
	ProjectID    uuid.UUID
}

type ProjectMetadatum struct {
	ProjectID uuid.UUID
	Metadata  []byte
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

type ProjectUser struct {
	ProjectID uuid.UUID
	UserID    uuid.UUID
}

type Role struct {
	RoleID          uuid.UUID
	Role            UserRoles
	RoleDescription pgtype.Text
}

type SystemLog struct {
	LogID     uuid.UUID
	Message   string
	Context   []byte
	CreatedAt pgtype.Timestamptz
	UserID    uuid.UUID
}

type User struct {
	UserID    uuid.UUID
	FirstName string
	LastName  string
	Password  string
	Email     string
	CreatedAt pgtype.Timestamp
	LastLogin pgtype.Timestamp
	RoleID    uuid.UUID
}