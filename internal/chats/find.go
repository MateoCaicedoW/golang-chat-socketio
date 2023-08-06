package chats

import (
	"database/sql"

	"github.com/gofrs/uuid"
)

func Find(tx *sql.DB, id uuid.UUID) (Chat, error) {
	var chat Chat
	err := tx.QueryRow(`
		SELECT 
			chats.id,
			users.id first_user_id, 
			users.name first_user_name,
			users.color first_user_color,
			users2.id second_user_id,
			users2.color second_user_color, 
			users2.name second_user_name
		FROM chats
		JOIN users ON chats.first_user = users.id
		JOIN users users2 ON chats.second_user = users2.id
		WHERE chats.id = $1;
	`, id).Scan(&chat.ID, &chat.FirstUserID, &chat.FirstUserName, &chat.FirstUserColor, &chat.SecondUserID, &chat.SecondUserColor, &chat.SecondUserName)
	if err != nil {
		return chat, err
	}

	return chat, nil
}
