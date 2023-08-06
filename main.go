package main

import (
	"log"
	"main/actions/chats"
	"main/actions/login"
	"main/chat"
	"main/middleware"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	server := chat.NewChatServer()
	go server.Serve()
	router.HandleFunc("/", login.New).Methods("GET")
	router.HandleFunc("/login/create", login.Create).Methods("POST")
	router.Handle("/socket.io/", server)

	routeProtected := router.PathPrefix("/").Subrouter()
	routeProtected.Use(middleware.IsLogged)
	routeProtected.HandleFunc("/chats", chats.List).Methods("GET")
	routeProtected.HandleFunc("/chats/new", chats.New).Methods("GET")
	routeProtected.HandleFunc("/chats/{user_id}/show", chats.Show).Methods("GET")

	srv := &http.Server{
		Handler:      router,
		Addr:         "localhost:3000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Serving at localhost:3000...")
	log.Fatal(srv.ListenAndServe())

}
