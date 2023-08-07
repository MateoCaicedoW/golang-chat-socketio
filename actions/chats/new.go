package chats

import (
	"main/db"
	"main/internal/users"
	"main/render"
	"main/session"
	"net/http"

	"github.com/gofrs/uuid"
)

func New(w http.ResponseWriter, r *http.Request) {
	currentUser := session.GetCurrentUser(r)
	userList, err := users.All(db.Tx, uuid.FromStringOrNil(currentUser))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.SetData("currentUser", currentUser)
	render.SetData("users", userList)
	render.RenderWithLayout(w, "/chats/new.html", "application.html")
}
