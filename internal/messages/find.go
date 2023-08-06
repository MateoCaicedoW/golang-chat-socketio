package messages

import (
	"database/sql"

	"github.com/gofrs/uuid"
)

func Find(tx *sql.DB, messageID uuid.UUID) (Message, error) {
	var message Message
	err := tx.QueryRow(`
	SELECT 
		messages.id,
		messages.chat_id, 
		content, 
		date, 
		users.email AS sender_email, 
		users.color AS sender_color,
		users.id AS sender_id, 
		users.name AS sender_name
	FROM messages
	JOIN users ON messages.sender_id = users.id
	WHERE messages.id = $1;
	`, messageID).Scan(&message.ID, &message.ChatID, &message.Content, &message.Date, &message.SenderEmail, &message.SenderColor, &message.SenderID, &message.SenderName)

	if err != nil {
		return message, err
	}

	message.SenderInital = message.SenderInitial()
	return message, nil
}
