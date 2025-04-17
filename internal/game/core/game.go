package core

import (
	"ashab-k/github.com/internal/game/entity"
	"ashab-k/github.com/internal/game/physics"
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth   = 320
	screenHeight  = 240
	ballRadius    = 5
	ballXSpeed    = 4
	ballYSpeed    = 4
	paddleSpeed   = 5
	paddleHeight  = 30
	paddleWidth   = 5
)

// Game represents the game state
type Game struct {
	Entities map[string]interface{}
	Ball     *entity.Ball
	Paddles  []*entity.Paddle
	FrameCount int
}

func (g *Game) Update() error {
	g.FrameCount++

	if g.FrameCount%180 == 0 { 
        fmt.Println("Game is running...")
    }
	// Update ball position
	g.Ball.Update(screenWidth, screenHeight)
	
	// Handle paddle input and update paddles
	g.handleInput()
	
	// Check for collisions with each paddle
	for _, paddle := range g.Paddles {
		physics.CheckPaddleBallCollision(g.Ball, paddle)
	}
	
	return nil
}

// Handle input for moving the paddles
func (g *Game) handleInput() {
	// Left paddle movement (W/S keys)
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.Paddles[0].MoveUp()
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.Paddles[0].MoveDown()
	}
	
	// Right paddle movement (Up/Down arrow keys)
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.Paddles[1].MoveUp()
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.Paddles[1].MoveDown()
	}

	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		fmt.Println("Game Quit")
        log.Fatal("Error")
    }
	
	// Update paddle positions with proper boundaries
	g.Paddles[0].Update(0, screenHeight-paddleHeight)
	g.Paddles[1].Update(0, screenHeight-paddleHeight)
}

func NewGame() *Game {
	
	ball := entity.NewBall(screenWidth/2, screenHeight/2, ballXSpeed, ballYSpeed, ballRadius)
	
	
	paddle1 := entity.NewPaddle(0, 0, paddleSpeed, 1, paddleHeight, paddleWidth)
	paddle2 := entity.NewPaddle(screenWidth-paddleWidth, screenHeight/2, paddleSpeed, 2, paddleHeight, paddleWidth)
	
	paddles := []*entity.Paddle{paddle1, paddle2}
	
	game := &Game{
		Ball:     ball,
		Paddles:  paddles,
		Entities: make(map[string]interface{}),
	}
	
	game.Entities["ball"] = ball
	game.Entities["paddles"] = paddles
	
	return game
}

func (g *Game) GetScreenDimensions() (width, height int) {
	return screenWidth, screenHeight
}