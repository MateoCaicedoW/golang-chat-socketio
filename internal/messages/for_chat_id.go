package messages

import (
	"database/sql"
	"time"

	"github.com/gobuffalo/nulls"
	"github.com/gofrs/uuid"
)

type Message struct {
	ID           uuid.UUID    `json:"id" db:"id"`
	ChatID       uuid.UUID    `json:"chat_id" db:"chat_id"`
	Content      string       `json:"content" db:"content"`
	Date         time.Time    `json:"date" db:"date"`
	SenderID     uuid.UUID    `json:"sender_id" db:"sender_id"`
	SenderName   string       `json:"sender_name" db:"sender_name"`
	SenderColor  nulls.String `json:"sender_color" db:"sender_color"`
	SenderEmail  string       `json:"sender_email" db:"sender_email"`
	SenderInital string       `json:"sender_initial" db:"-"`
}

func ForChatID(tx *sql.DB, chatID uuid.UUID) ([]Message, error) {
	rows, err := tx.Query(`
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
	WHERE messages.chat_id = $1
	ORDER BY messages.date ASC;
	`, chatID)

	if err != nil {
		return nil, err
	}

	var messages []Message
	for rows.Next() {
		var message Message
		if err := rows.Scan(&message.ID, &message.ChatID, &message.Content, &message.Date, &message.SenderEmail, &message.SenderColor, &message.SenderID, &message.SenderName); err != nil {
			return nil, err
		}

		messages = append(messages, message)
	}

	return messages, nil
}

func (message Message) SenderInitial() string {
	if message.SenderName != "" {
		return message.SenderName[:1]
	}

	return ""
}
