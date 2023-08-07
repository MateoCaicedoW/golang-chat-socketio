package chats

import (
	"main/db"
	"main/internal/users"
	"main/render"
	"net/http"
)

func New(w http.ResponseWriter, r *http.Request) {
	currentUser := render.GetData("currentUser").(users.User)
	userList, err := users.All(db.Tx, currentUser.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.SetData("users", userList)
	render.RenderWithLayout(w, "/chats/new.html", "application.html")
}
