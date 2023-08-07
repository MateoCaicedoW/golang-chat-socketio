package middleware

import (
	"main/db"
	"main/internal/users"
	"main/render"
	"main/session"
	"net/http"

	"github.com/gofrs/uuid"
)

func SetCurrentUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//set current user
		currentUser := session.GetCurrentUser(r)

		user, err := users.Find(db.Tx, uuid.FromStringOrNil(currentUser))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		render.SetData("currentUser", user)

		next.ServeHTTP(w, r)
	})
}
