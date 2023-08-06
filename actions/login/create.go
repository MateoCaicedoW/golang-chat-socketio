package login

import (
	"main/db"
	"main/models"
	"main/session"
	"net/http"

	"github.com/go-playground/form/v4"
)

func Create(w http.ResponseWriter, r *http.Request) {
	isLogged := session.IsLoggedIn(r)
	if isLogged {
		http.Redirect(w, r, "/chats", http.StatusFound)
		return
	}

	decoder := form.NewDecoder()
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var user models.User
	if err := decoder.Decode(&user, r.Form); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result := db.Tx.QueryRow("SELECT * FROM users WHERE email = $1", user.Email)
	if err := result.Scan(&user.ID, &user.Email, &user.Name, &user.Color); err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	session.SetCurrentUser(w, r, user)

	http.Redirect(w, r, "/chats", http.StatusFound)
}
