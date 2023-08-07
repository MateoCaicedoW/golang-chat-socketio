package chats

import (
	"main/db"
	"main/internal/chats"
	"main/internal/users"
	"main/render"
	"net/http"
)

func List(w http.ResponseWriter, r *http.Request) {
	currentUser := render.GetData("currentUser").(users.User)
	chatList, err := chats.ListForUserID(db.Tx, currentUser.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.SetData("chats", chatList)
	render.RenderWithLayout(w, "/chats/list.html", "application.html")
}
