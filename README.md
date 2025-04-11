# LivechatRoom
Live Chat Room in Golang 

1.SetUp
  1. Git clone : git clone https://github.com/<<username>>/LivechatRoom.git
  2. cd LivechatRoom
2. Initialize the Go Module
    Make sure you have Go installed (v1.20 or later).
     go mod tidy
3. Run the Server
    go run main.go
   ![runningLivechat](https://github.com/user-attachments/assets/b08f5dbe-0bf4-41cf-8eb0-013985ccebcd)

4. API Endpoints
🟢 Join Chat
       Method: GET

       Endpoint: /join?id=<client_id>

       Example: /join?id=alice

Adds a new client to the chat room.

✉️ Send Message
       Method: GET

       Endpoint: /send?id=<client_id>&message=<your_message>

       Example: /send?id=alice&message=Hello everyone!

Broadcasts a message to all connected clients.

📬 Receive Messages
      Method: GET

      Endpoint: /messages?id=<client_id>

      Behavior: Waits for a new message for up to 10 seconds, then times out.

      Example: /messages?id=alice

❌ Leave Chat
      Method: GET

      Endpoint: /leave?id=<client_id>

      Example: /leave?id=alice

Unsubscribes the client from the chat room.

5. 🛠 Tech Stack
       💬 Go (v1.20+)

       🧵 Goroutines & Channels

      🌐 net/http standard library (no frameworks)

      🧪 cURL / YARC / Postman for testing

6.Demo vedios


