package models

import (
	"github.com/gobuffalo/nulls"
	"github.com/gofrs/uuid"
)

type User struct {
	ID    uuid.UUID    `json:"id" db:"id"`
	Name  string       `json:"name" db:"name"`
	Email string       `json:"email" db:"email"`
	Color nulls.String `json:"color" db:"color"`
}

type Users []User
