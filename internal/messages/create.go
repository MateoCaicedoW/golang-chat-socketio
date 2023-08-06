package messages

import (
	"database/sql"

	"github.com/gofrs/uuid"
)

func Create(tx *sql.DB, senderID, chatID uuid.UUID, content string) (Message, error) {

	var id uuid.UUID
	err := tx.QueryRow(
		`INSERT INTO messages (id, chat_id, content, sender_id, date)
		VALUES ($1, $2, $3, $4, NOW())
		RETURNING id;
		`,
		uuid.Must(uuid.NewV4()), chatID, content, senderID).Scan(&id)
	if err != nil {
		return Message{}, err
	}

	msg, err := Find(tx, id)
	if err != nil {
		return Message{}, err
	}

	return msg, nil

}
