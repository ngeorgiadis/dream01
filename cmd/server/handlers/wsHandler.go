package handlers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"dream01/internal/intlog"

	"github.com/gorilla/websocket"
)

// Message ...
type Message struct {
	ClientID    string `json:"clientId"`
	MessageType string `json:"messageType"`
}

var broadcast = make(chan Message)
var clients = make(map[*websocket.Conn]bool)
var upgrader = websocket.Upgrader{}
var mutex = &sync.Mutex{}

// HandleMessages ...
func HandleMessages() {

	for {

		msg := <-broadcast

		for client := range clients {
			mutex.Lock()
			err := client.WriteJSON(msg)
			if err != nil {
				intlog.Errorf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
			mutex.Unlock()
		}

	}

}

func wsHandler(w http.ResponseWriter, r *http.Request) {

	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Make sure we close the connection when the function returns
	defer ws.Close()

	// Register our new client
	//clients[ws] = true

	for {
		var msg Message

		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			//delete(clients, ws)
			break
		}

		switch msg.MessageType {
		case "ping":
			var pong Message
			pong.MessageType = "pong"
			pong.ClientID = msg.ClientID
			ws.WriteJSON(pong)
		default:
			//broadcast <- msg
		}

		// Send the newly received message to the broadcast channel

	}
}

func pingClients() {
	for {

		for client := range clients {

			var msg Message

			msg.MessageType = "ping"
			err := client.WriteJSON(msg)

			if err != nil {
				intlog.Errorf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}

		time.Sleep(15 * time.Second)
	}
}
