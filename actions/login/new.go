package login

import (
	"main/render"
	"main/session"
	"net/http"
)

func New(w http.ResponseWriter, r *http.Request) {
	isLogged := session.IsLoggedIn(r)
	if isLogged {
		http.Redirect(w, r, "/chats", http.StatusFound)
		return
	}

	render.RenderWithLayout(w, "/login/login.html", "application.html")
}
