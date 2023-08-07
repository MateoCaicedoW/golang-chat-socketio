package login

import (
	"main/session"
	"net/http"
)

func Destroy(w http.ResponseWriter, r *http.Request) {
	session.Destroy(w, r)
	http.Redirect(w, r, "/", http.StatusFound)
}
