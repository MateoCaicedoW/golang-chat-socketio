package models

import "github.com/gofrs/uuid"

type Chat struct {
	ID           int       `json:"id" db:"id"`
	FirstUserID  uuid.UUID `json:"first_user" db:"first_user"`
	SecondUserID uuid.UUID `json:"second_user" db:"second_user"`
}

type Chats []Chat
