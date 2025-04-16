package entity

type Paddle struct {
	X      float32
	Y      float32
	Dx     float32 // Horizontal speed
	Dy     float32 // Vertical speed
	Height float32
	Width  float32
}

// MoveUp moves the paddle up
func (p *Paddle) MoveUp() {
	p.Y -= p.Dy // Move up by subtracting vertical speed
}

// MoveDown moves the paddle down
func (p *Paddle) MoveDown() {
	p.Y += p.Dy // Move down by adding vertical speed
}

// Update constrains the paddle to the screen boundaries
func (p *Paddle) Update(minY, maxY float32) {
	// Keep paddle within vertical bounds
	if p.Y < minY {
		p.Y = minY // Set to minimum Y if too high
	}
	if p.Y > maxY {
		p.Y = maxY // Set to maximum Y if too low
	}
}

func NewPaddle(x, y, dy, dx, paddleHeight, paddleWidth float32) *Paddle {
	return &Paddle{
		X:      x,      
		Y:      y,    
		Dy:     dy,    
		Dx:     dx,     
		Height: paddleHeight,
		Width:  paddleWidth,
	}
}