package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var message string

type requestBody struct {
	Message string `json:"message"`
}

// get
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	var messages []Message

	result := DB.Find(&messages)
	if result.Error != nil {
		http.Error(w, "Failed to fetch messages", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(messages)
	if err != nil {
		http.Error(w, "Failed to encode messages", http.StatusInternalServerError)
		return
	}
}

// create
func CreateMessageHandler(w http.ResponseWriter, r *http.Request) {
	var reqBody requestBody

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, "Invalid request Body", http.StatusBadRequest)
		return
	}

	message = reqBody.Message

	newMessage := Message{
		Text: message,
	}

	result := DB.Create(&newMessage)
	if result.Error != nil {
		http.Error(w, "Failed to save message", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Message updated and saved successfully")
}

// update
func UpdateMessageByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var message Message
	result := DB.First(&message, id)
	if result.Error != nil {
		http.Error(w, "Message not found", http.StatusNotFound)
		return
	}

	var reqBody requestBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	message.Text = reqBody.Message
	DB.Save(&message)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Message with ID %s updated successfully", id)
}

// delete
func DeleteMessageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var message Message
	result := DB.First(&message, id)
	if result.Error != nil {
		http.Error(w, "Message not found", http.StatusNotFound)
	}

	DB.Delete(&message)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Message with ID %s deleted successfully", id)
}

func main() {
	InitDB()
	DB.AutoMigrate(&Message{})

	router := mux.NewRouter()

	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
	router.HandleFunc("/api/message", CreateMessageHandler).Methods("POST")
	router.HandleFunc("/api/message/{id}", UpdateMessageByIDHandler).Methods("PATCH")
	router.HandleFunc("/api/message/{id}", DeleteMessageHandler).Methods("DELETE")

	http.ListenAndServe(":8080", router)
}
