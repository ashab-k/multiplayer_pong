package main

import (
	"ashab-k/github.com/internal/game/core"
	"ashab-k/github.com/internal/render"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
    screenWidth  = 320
    screenHeight = 240
    windowWidth  = 640
    windowHeight = 480
)

// GameWrapper implements ebiten.Game interface
type GameWrapper struct {
    game     *core.Game
    renderer *render.Renderer
}

func (g *GameWrapper) Update() error {
    return g.game.Update()
}

func (g *GameWrapper) Draw(screen *ebiten.Image) {
    g.renderer.Render(screen, g.game)
}

func (g *GameWrapper) Layout(outsideWidth, outsideHeight int) (int, int) {
    return screenWidth, screenHeight
}

func main() {
    // Create game and renderer
    game := core.NewGame()
    renderer := render.NewRenderer()
    
    // Create wrapper that connects game state to rendering
    wrapper := &GameWrapper{
        game:     game,
        renderer: renderer,
    }
    
    // Set up window
    ebiten.SetWindowSize(windowWidth, windowHeight)
    ebiten.SetWindowTitle("Pong")
    
    // Start game loop
    if err := ebiten.RunGame(wrapper); err != nil {
        log.Fatal(err)
    }
}