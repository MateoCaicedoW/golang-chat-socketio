package chats

import (
	"database/sql"

	"github.com/gofrs/uuid"
)

func Exists(tx *sql.DB, firstUserID uuid.UUID, secondUserID uuid.UUID) (uuid.UUID, error) {
	var id uuid.UUID
	err := tx.QueryRow(`
	SELECT chats.id FROM chats
	WHERE (chats.first_user = $1 OR chats.second_user = $1) 
	AND  (chats.first_user = $2 OR chats.second_user = $2)
	LIMIT 1;
	`, firstUserID, secondUserID).Scan(&id)

	if err != nil {
		return uuid.Nil, nil
	}

	return id, nil
}
