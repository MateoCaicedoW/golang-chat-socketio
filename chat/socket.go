package chat

import (
	"fmt"

	socketio "github.com/googollee/go-socket.io"
)

var server = socketio.NewServer(nil)

// implement the chat server which will be used in main.go only to start the server
func NewChatServer() *socketio.Server {
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		return nil
	})

	// write a function to hanlde the join room event
	server.OnEvent("/", "join", func(s socketio.Conn, room string) {
		fmt.Println("join")
		fmt.Println("room: ", room)
		s.Join(room)
		s.Emit("join", room)

	})

	setMethods()
	return server
}

func setMethods() *socketio.Server {
	// server.JoinRoom("/", "chat")

	server.OnEvent("/chat", "msg", SendMessage)

	return server
}
