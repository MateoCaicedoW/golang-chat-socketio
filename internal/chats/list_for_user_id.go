package chats

import (
	"database/sql"

	"github.com/gobuffalo/nulls"
	"github.com/gofrs/uuid"
)

type Chat struct {
	ID           uuid.UUID `json:"id" db:"id"`
	FirstUserID  uuid.UUID `json:"first_user" db:"first_user_id"`
	SecondUserID uuid.UUID `json:"second_user" db:"second_user_id"`

	FirstUserName   string       `json:"first_user_name" db:"first_user_name"`
	FirstUserColor  nulls.String `json:"first_user_color" db:"first_user_color"`
	SecondUserName  string       `json:"second_user_name" db:"second_user_name"`
	SecondUserColor nulls.String `json:"second_user_color" db:"second_user_color"`
}

func ListForUserID(tx *sql.DB, userID uuid.UUID) ([]Chat, error) {
	rows, err := tx.Query(`
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
	WHERE chats.first_user = $1 OR chats.second_user = $1;

	`, userID)
	if err != nil {
		return nil, err
	}

	var chats []Chat
	for rows.Next() {
		var chat Chat
		if err := rows.Scan(&chat.ID, &chat.FirstUserID, &chat.FirstUserName, &chat.FirstUserColor, &chat.SecondUserID, &chat.SecondUserColor, &chat.SecondUserName); err != nil {
			return nil, err
		}
		chats = append(chats, chat)
	}

	return chats, nil

}

func (chat Chat) SecondUserInitial() string {
	if chat.SecondUserName != "" {
		return string(chat.SecondUserName[0])
	}
	return "?"
}

func (chat Chat) FirstUserInitial() string {
	if chat.FirstUserName != "" {
		return string(chat.FirstUserName[0])
	}
	return "?"
}
