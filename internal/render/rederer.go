package render

import (
	"ashab-k/github.com/internal/game/core"
	"ashab-k/github.com/internal/game/entity"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// Renderer handles all game rendering
type Renderer struct {
}

func (r *Renderer) DrawBall(screen *ebiten.Image, ball *entity.Ball) {
    vector.DrawFilledCircle(screen, ball.X, ball.Y, ball.Radius, color.White, false)
}
func (r *Renderer) DrawPaddle(screen *ebiten.Image , paddle *entity.Paddle){
	vector.DrawFilledRect(screen  , paddle.X , paddle.Y , paddle.Width , paddle.Height , color.White , false)
}

// Render renders the entire game state
func (r *Renderer) Render(screen *ebiten.Image, game *core.Game) {
    r.DrawBall(screen, game.Ball)
	for _ , paddle := range game.Paddles {
		r.DrawPaddle(screen , paddle)
	}
    
}

func NewRenderer() *Renderer {
    return &Renderer{}
}

