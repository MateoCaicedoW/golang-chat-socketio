package chats

import (
	"database/sql"

	"github.com/gofrs/uuid"
)

func Create(tx *sql.DB, firstUserID uuid.UUID, secondUserID uuid.UUID) (Chat, error) {
	var id uuid.UUID
	err := tx.QueryRow(`
	INSERT INTO chats (id, first_user, second_user)
	VALUES ($1, $2, $3)
	RETURNING id;
	`, uuid.Must(uuid.NewV4()), firstUserID, secondUserID).Scan(&id)
	if err != nil {
		return Chat{}, err
	}

	chat, err := Find(tx, id)
	if err != nil {
		return Chat{}, err
	}

	return chat, nil
}
