package chat

import (
	"main/db"
	"main/internal/chats"

	"github.com/gofrs/uuid"
	socketio "github.com/googollee/go-socket.io"
)

func JoinRoom(s socketio.Conn, room map[string]string) {
	firstUserID := uuid.FromStringOrNil(room["firstUserID"])
	secondUserID := uuid.FromStringOrNil(room["secondUserID"])

	chatID, err := chats.Exists(db.Tx, firstUserID, secondUserID)
	if err != nil {
		return
	}

	if !chatID.IsNil() {
		s.Join(chatID.String())
		return
	}

	chat, err := chats.Create(db.Tx, firstUserID, secondUserID)
	if err != nil {
		return
	}

	roomName := chat.ID.String()

	s.Join(roomName)

}
