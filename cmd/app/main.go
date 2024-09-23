package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"web-app/internal/database"
	"web-app/internal/handlers"
	"web-app/internal/messagesService"
)

func main() {
	database.InitDB()
	database.DB.AutoMigrate(&messagesService.Message{})

	repo := messagesService.NewMessageRepository(database.DB)
	service := messagesService.NewService(repo)
	handler := handlers.NewHandler(service)

	router := mux.NewRouter()
	router.HandleFunc("/api/hello", handler.GetMessagesHandler).Methods("GET")
	router.HandleFunc("/api/message", handler.PostMessageHandler).Methods("POST")
	router.HandleFunc("/api/message/{id}", handler.UpdateMessageByIDHandler).Methods("PATCH")
	router.HandleFunc("/api/message/{id}", handler.DeleteMessageByIDHandler).Methods("DELETE")

	http.ListenAndServe(":8080", router)
}
