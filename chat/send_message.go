package chat

import (
	"fmt"
	"main/db"
	"main/internal/chats"
	"main/internal/messages"

	"github.com/gofrs/uuid"
	socketio "github.com/googollee/go-socket.io"
)

func SendMessage(s socketio.Conn, message map[string]string) string {
	firstUserID := uuid.FromStringOrNil(message["receiverID"])
	secondUserID := uuid.FromStringOrNil(message["senderID"])

	chatID, err := chats.Exists(db.Tx, firstUserID, secondUserID)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	msg, err := messages.Create(db.Tx, uuid.FromStringOrNil(message["senderID"]), chatID, message["message"])
	if err != nil {
		fmt.Println(err)
		return ""
	}

	message = map[string]string{
		"content":        msg.Content,
		"sender_id":      msg.SenderID.String(),
		"chat_id":        msg.ChatID.String(),
		"date":           msg.Date.String(),
		"sender_color":   msg.SenderColor.String,
		"name":           msg.SenderName,
		"email":          msg.SenderEmail,
		"sender_initial": msg.SenderInitial(),
		"receiver_id":    message["receiverID"],
	}

	server.BroadcastToRoom("/", chatID.String(), "reply", message)
	return "recv " + message["msg"]
}
