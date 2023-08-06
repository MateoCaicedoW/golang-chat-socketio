package chats

import (
	"main/db"
	"main/internal/chats"
	"main/internal/messages"
	"main/internal/users"
	"main/render"
	"main/session"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"
)

func Show(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID := params["user_id"]
	currentUser := session.GetCurrentUser(r)
	id, err := chats.Exists(db.Tx, uuid.FromStringOrNil(currentUser), uuid.FromStringOrNil(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user, err := users.Find(db.Tx, uuid.FromStringOrNil(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userLogged, err := users.Find(db.Tx, uuid.FromStringOrNil(currentUser))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.SetData("chatID", id)
	render.SetData("user", user)
	render.SetData("userLogged", userLogged)
	render.SetData("currentUser", uuid.FromStringOrNil(currentUser))
	//set the messages only if the chat exists
	messagesList := []messages.Message{}
	if id.IsNil() {
		render.SetData("messages", messagesList)
		render.RenderWithLayout(w, "/chats/show.html", "application.html")
		return
	}

	messagesList, err = messages.ForChatID(db.Tx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.SetData("messages", messagesList)
	render.RenderWithLayout(w, "/chats/show.html", "application.html")
}
