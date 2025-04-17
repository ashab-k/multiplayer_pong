package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{} 

type Player struct {
	conn *websocket.Conn
}

var players []*Player

func HandleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade failed:", err)
		return
	}

	player := &Player{conn: conn}
	players = append(players, player)

	defer func() {
		conn.Close()
	}()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}
		log.Println("Received from client:", string(msg))

		// Echo to all other players
		for _, p := range players {
			if p != player {
				p.conn.WriteMessage(websocket.TextMessage, msg)
			}
		}
	}
}

func main() {
	http.HandleFunc("/ws", HandleConnection)
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
