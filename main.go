package main

import (
	"LivechatRoom/chat"
	"LivechatRoom/handlers"
	"log"
	"net/http"
)

func main() {
	chatRoom := chat.NewChatRoom()

	http.HandleFunc("/join", handlers.JoinHandler(chatRoom))
	http.HandleFunc("/send", handlers.SendHandler(chatRoom))
	http.HandleFunc("/leave", handlers.LeaveHandler(chatRoom))
	http.HandleFunc("/messages", handlers.MessagesHandler(chatRoom))

	log.Println("Server started at :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
