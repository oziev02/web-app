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

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", message)
}

func UpdateMessageHandler(w http.ResponseWriter, r *http.Request) {
	var reqBody requestBody

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, "Invalid request Body", http.StatusBadRequest)
		return
	}

	message = reqBody.Message
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
	router.HandleFunc("/api/message", UpdateMessageHandler).Methods("POST")
	http.ListenAndServe(":8080", router)
}
