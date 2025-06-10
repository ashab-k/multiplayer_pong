package main

import (
	"ashab-k/github.com/internal/game/core"
	"ashab-k/github.com/internal/render"
	"encoding/json"
	"io"
	"log"

	"github.com/gorilla/websocket"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
    screenWidth  = 320
    screenHeight = 240
    windowWidth  = 640
    windowHeight = 480
)

type BallState struct {
	X  float32 `json:"x"`
	Y  float32 `json:"y"`
	Dx float32 `json:"dx"`
	Dy float32 `json:"dy"`
}

func Connect()(*websocket.Conn , error){
    conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
	if err != nil {
		log.Fatal("Dial error: ", err)
	}
    return conn , err
}
// GameWrapper implements ebiten.Game interface
type GameWrapper struct {
    game     *core.Game
    renderer *render.Renderer
    ballStateChan chan BallState
}

func (g *GameWrapper) Update() error {

     select {
    case ballState := <-g.ballStateChan:
        // Update your game's ball position with server state
        // You'll need to add a method to your core.Game to update ball state
        g.game.UpdateBallFromServer(ballState.X, ballState.Y, ballState.Dx, ballState.Dy)
    default:
        // No update from server
    }
    return g.game.Update()
}

func (g *GameWrapper) Draw(screen *ebiten.Image) {
    g.renderer.Render(screen, g.game)
}

func (g *GameWrapper) Layout(outsideWidth, outsideHeight int) (int, int) {
    return screenWidth, screenHeight
}


func main() {
    IsSinglePlayer := false // Change this since you're connecting to server
    conn, err := Connect()
    if err != nil {
        log.Fatal("Error in establishing connection with websocket")
    }
    
    game := core.NewGame(IsSinglePlayer)
    renderer := render.NewRenderer()
    
    // Channel to communicate ball state updates
    ballStateChan := make(chan BallState, 1)
    
    go func() {
        for {
            _, r, err := conn.NextReader()
            if err != nil {
                log.Println("Read error:", err)
                return
            }
            message, err := io.ReadAll(r)
            if err != nil {
                log.Println("ReadAll error:", err)
                continue
            }
            
            var state BallState
            err = json.Unmarshal(message, &state)
            if err != nil {
                log.Println("Failed to parse ball state:", err)
                continue
            }
            
            // Send ball state to game
            select {
            case ballStateChan <- state:
            default:
                // Channel is full, skip this update
            }
        }
    }()
    
    // Create wrapper that connects game state to rendering
    wrapper := &GameWrapper{
        game:         game,
        renderer:     renderer,
        ballStateChan: ballStateChan,
    }
    
    ebiten.SetWindowSize(windowWidth, windowHeight)
    ebiten.SetWindowTitle("Pong")
    
    if err := ebiten.RunGame(wrapper); err != nil {
        log.Fatal(err)
    }
}