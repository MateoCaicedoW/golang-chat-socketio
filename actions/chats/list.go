package chats

import (
	"main/db"
	"main/internal/chats"
	"main/render"
	"main/session"
	"net/http"

	"github.com/gofrs/uuid"
)

func List(w http.ResponseWriter, r *http.Request) {
	currentUser := session.GetCurrentUser(r)
	chatList, err := chats.ListForUserID(db.Tx, uuid.FromStringOrNil(currentUser))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.SetData("currentUser", currentUser)
	render.SetData("chats", chatList)
	render.RenderWithLayout(w, "/chats/list.html", "application.html")
}
