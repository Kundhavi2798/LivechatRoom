package chat

import (
	"fmt"
	"log"
)

type Client struct {
	ID      string
	Message chan string
}

type Message struct {
	SenderID string
	Content  string
}

type ChatRoom struct {
	Clients       map[string]*Client
	Register      chan *Client
	Unregister    chan *Client
	ClientMessage chan Message
}

func NewChatRoom() *ChatRoom {
	room := &ChatRoom{
		Clients:       make(map[string]*Client),
		Register:      make(chan *Client),
		Unregister:    make(chan *Client),
		ClientMessage: make(chan Message),
	}

	go room.run()
	return room
}

func (c *ChatRoom) run() {
	for {
		select {
		case client := <-c.Register:
			c.Clients[client.ID] = client
			log.Printf("Client %s joined\n", client.ID)

		case client := <-c.Unregister:
			if _, ok := c.Clients[client.ID]; ok {
				delete(c.Clients, client.ID)
				close(client.Message)
				log.Printf("Client %s left\n", client.ID)
			}

		case msg := <-c.ClientMessage:
			broadcast := fmt.Sprintf("%s: %s", msg.SenderID, msg.Content)
			for _, client := range c.Clients {
				select {
				case client.Message <- broadcast:
				default:
					// Drop message if buffer full
				}
			}
		}
	}
}
