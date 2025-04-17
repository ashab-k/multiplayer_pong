package entity

type Paddle struct {
	X      float32
	Y      float32
	Dx     float32 
	Dy     float32 
	Height float32
	Width  float32
}

// MoveUp moves the paddle up
func (p *Paddle) MoveUp() {
	p.Y -= p.Dy 
}

// MoveDown moves the paddle down
func (p *Paddle) MoveDown() {
	p.Y += p.Dy 
}

func (p *Paddle) Update(minY, maxY float32) {
	
	if p.Y < minY {
		p.Y = minY 
	}
	if p.Y > maxY {
		p.Y = maxY
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