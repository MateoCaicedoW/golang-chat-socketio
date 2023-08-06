package session

import (
	"main/models"
	"net/http"

	"github.com/gorilla/sessions"
)

var Store = sessions.NewCookieStore([]byte("session"))

// each session contains the username of the user and the time at which it expires
type Session struct {
	Username string
}

func SetCurrentUser(w http.ResponseWriter, r *http.Request, user models.User) {
	session, _ := Store.Get(r, "session")
	session.Values["user_id"] = user.ID.String()
	session.Save(r, w)
}

func GetCurrentUser(r *http.Request) string {
	session, _ := Store.Get(r, "session")
	return session.Values["user_id"].(string)
}

func IsLoggedIn(r *http.Request) bool {
	session, _ := Store.Get(r, "session")
	return session.Values["user_id"] != nil
}
