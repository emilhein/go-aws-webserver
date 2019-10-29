package webserver

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/websocket"

	"github.com/gorilla/mux"
)

var clients = make(map[*websocket.Conn]bool) // connected clients
var broadcast = make(chan Event)             // broadcast channel

// Configure the upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Define our message object
type Event struct {
	Type      string `json:"type"`
	Timestamp string `json:"timestamp"`
	Content   string `json:"content"`
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Make sure we close the connection when the function returns
	defer ws.Close()

	// Register our new client
	clients[ws] = true

	for {
		var msg Event
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}
		fmt.Println("We recieved start signal!")
		go startProcess(broadcast)

		// Send the newly received message to the broadcast channel
		broadcast <- msg
	}
}

func handleMessages() {
	limiter := time.Tick(time.Millisecond * 200)

	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast
		<-limiter
		// Send it out to every client that is currently connected
		for client := range clients {

			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func Start() {
	r := mux.NewRouter()
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	r.HandleFunc("/ws", handleConnections)
	go handleMessages()

	r.HandleFunc("/simple", Simple)
	r.HandleFunc("/getS3files", GetS3Files).Methods("POST")
	r.HandleFunc("/interfaces", InterfaceMethod)
	r.HandleFunc("/startmining", StartMining)
	fmt.Printf("Server started on port :%v \n", port)
	portString := fmt.Sprintf(":%v", port)
	err := http.ListenAndServe(portString, r)
	if err != nil {
		fmt.Printf("Could not start the server: %v", err)
	}

}
