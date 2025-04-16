package entity

// Ball represents a ball entity in the game world
type Ball struct {
    X, Y    float32  
    Dx, Dy  float32  
    Radius  float32  
}

func (b *Ball) Update(screenWidth, screenHeight float32) {
    b.X += b.Dx
    b.Y += b.Dy
    
    if b.X >= screenWidth || b.X <= 0 {
        b.Dx = -b.Dx
    }
    
    if b.Y >= screenHeight || b.Y <= 0 {
        b.Dy = -b.Dy
    }
}

func NewBall(x, y, dx, dy, radius float32) *Ball {
    return &Ball{
        X: x,
        Y: y,
        Dx: dx,
        Dy: dy,
        Radius: radius,
    }
}