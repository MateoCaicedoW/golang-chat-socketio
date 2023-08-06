package middleware

import (
	"main/session"
	"net/http"
)

func IsLogged(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//check if user is logged in
		isLogged := session.IsLoggedIn(r)
		if !isLogged {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

		next.ServeHTTP(w, r)
	})
}
