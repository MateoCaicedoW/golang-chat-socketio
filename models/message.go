package models

type Message struct {
	ID             int    `json:"id" db:"id"`
	ConversationID int    `json:"conversation_id" db:"conversation_id"`
	SenderID       int    `json:"sender_id" db:"sender_id"`
	Content        string `json:"content" db:"content"`
	Date           string `json:"created_at" db:"created_at"`
}

type Messages []Message
