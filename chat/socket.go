package chat

import (
	"fmt"

	socketio "github.com/googollee/go-socket.io"
)

var server = socketio.NewServer(nil)

// implement the chat server which will be used in main.go only to start the server
func NewChatServer() *socketio.Server {

	server.OnEvent("/", "msg", SendMessage)
	server.OnEvent("/", "join", JoinRoom)
	server.OnEvent("/", "leave", LeaveRoom)
	server.OnEvent("/", "chat", SendMessage)
	server.OnEvent("/", "typing", Typing)
	server.OnEvent("/", "stop typing", StopTyping)
	server.OnEvent("/", "disconnect", Disconnect)
	return server
}

func LeaveRoom(s socketio.Conn, room string) {
	s.Leave(room)
	fmt.Println("leave", room)
}

func Typing(s socketio.Conn, room string) {
	fmt.Println("typing", room)
	server.BroadcastToRoom("/", room, "typing", "")
}

func StopTyping(s socketio.Conn, room string) {
	fmt.Println("stop typing", room)
	server.BroadcastToRoom("/", room, "stop typing", "")
}

func Disconnect(s socketio.Conn, room string) {
	fmt.Println("disconnect", room)
	server.BroadcastToRoom("/", room, "disconnect", "")
}
