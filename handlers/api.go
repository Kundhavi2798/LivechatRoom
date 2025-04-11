package handlers

import (
	"LivechatRoom/chat"
	"fmt"
	"net/http"
	"time"
)

func JoinHandler(chatRoom *chat.ChatRoom) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "Missing id", http.StatusBadRequest)
			return
		}
		client := &chat.Client{
			ID:      id,
			Message: make(chan string, 10),
		}
		chatRoom.Register <- client
		fmt.Fprintf(w, "Client %s joined.\n", id)
	}
}

func SendHandler(chatRoom *chat.ChatRoom) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		message := r.URL.Query().Get("message")

		if id == "" || message == "" {
			http.Error(w, "Missing id or message", http.StatusBadRequest)
			return
		}

		chatRoom.ClientMessage <- chat.Message{SenderID: id, Content: message}
		fmt.Fprintf(w, "Message sent.\n")
	}
}

func LeaveHandler(chatRoom *chat.ChatRoom) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "Missing id", http.StatusBadRequest)
			return
		}

		if client, ok := chatRoom.Clients[id]; ok {
			chatRoom.Unregister <- client
			fmt.Fprintf(w, "Client %s left.\n", id)
		} else {
			http.Error(w, "Client not found", http.StatusNotFound)
		}
	}
}

func MessagesHandler(chatRoom *chat.ChatRoom) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "Missing id", http.StatusBadRequest)
			return
		}

		client, ok := chatRoom.Clients[id]
		if !ok {
			http.Error(w, "Client not found", http.StatusNotFound)
			return
		}

		select {
		case msg := <-client.Message:
			fmt.Fprint(w, msg)
		case <-time.After(10 * time.Second):
			http.Error(w, "Timeout waiting for message", http.StatusRequestTimeout)
		}
	}
}
