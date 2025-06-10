package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{} 
type GameRoom struct {
    players []*Player
    ball    BallState
}

var gameRooms []*GameRoom

type Player struct {
    conn *websocket.Conn
    room *GameRoom
}

var players []*Player


type BallState struct {
	X  float32 `json:"x"`
	Y  float32 `json:"y"`
	Dx float32 `json:"dx"`
	Dy float32 `json:"dy"`
}

type GameState struct  { 
	Ball BallState
}

func HandleConnection(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println("Upgrade failed:", err)
        return
    }
    
    player := &Player{conn: conn}
    
    // Find or create a game room
    var room *GameRoom
    for _, r := range gameRooms {
        if len(r.players) < 2 { // Max 2 players per room
            room = r
            break
        }
    }
    
    if room == nil {
        // Create new room
        room = &GameRoom{
            players: []*Player{},
            ball:    BallState{X: 160, Y: 120, Dx: 4, Dy: 4},
        }
        gameRooms = append(gameRooms, room)
    }
    
    room.players = append(room.players, player)
    player.room = room
    
    defer func() {
        // Remove player from room when disconnecting
        for i, p := range room.players {
            if p == player {
                room.players = append(room.players[:i], room.players[i+1:]...)
                break
            }
        }
        conn.Close()
    }()
    
    for {
        _, msg, err := conn.ReadMessage()
        if err != nil {
            log.Println("Read error:", err)
            break
        }
        
        // Handle player input (paddle movement, etc.)
        // Broadcast to other players in the same room
        for _, p := range room.players {
            if p != player {
                p.conn.WriteMessage(websocket.TextMessage, msg)
            }
        }
    }
}


func gameLoop() {
    for {
        time.Sleep(time.Millisecond * 16)
        
        // Update physics for each active game room
        for _, room := range gameRooms {
            if len(room.players) > 0 { // Only update if there are players
                // Physics
                room.ball.X += room.ball.Dx
                room.ball.Y += room.ball.Dy
                
                if room.ball.X <= 0 || room.ball.X >= 320 {
                    room.ball.Dx *= -1
                }
                if room.ball.Y <= 0 || room.ball.Y >= 240 {
                    room.ball.Dy *= -1
                }
                
                // Broadcast to all players in this room
                data, _ := json.Marshal(room.ball)
                for _, p := range room.players {
                    err := p.conn.WriteMessage(websocket.TextMessage, data)
                    if err != nil {
                        log.Println("Write error:", err)
                    }
                }
            }
        }
    }
}


func main() {
	http.HandleFunc("/ws", HandleConnection)
	go gameLoop()
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
