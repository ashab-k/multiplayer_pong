package physics

import (
	"ashab-k/github.com/internal/game/entity"
)

func CheckPaddleBallCollision(ball *entity.Ball, paddle *entity.Paddle) bool {
    // Calculate paddle boundaries
    paddleLeft := paddle.X - paddle.Width/2
    paddleRight := paddle.X + paddle.Width/2
    paddleTop := paddle.Y - paddle.Height/2
    paddleBottom := paddle.Y + paddle.Height/2
    
    if ball.X+ball.Radius >= paddleLeft && 
       ball.X-ball.Radius <= paddleRight &&
       ball.Y+ball.Radius >= paddleTop &&
       ball.Y-ball.Radius <= paddleBottom {
        
        handlePaddleBallResponse(ball, paddle)
        return true
    }
    
    return false
}

func handlePaddleBallResponse(ball *entity.Ball, paddle *entity.Paddle) {
    ball.Dx = -ball.Dx 
    
    hitPosition := (ball.Y - paddle.Y) / (paddle.Height / 2)
    ball.Dy = hitPosition * 2
    
   if paddle.X < 160 {
	ball.X = paddle.X + paddle.Width/2 + ball.Radius + 1
   }else {
	ball.X = paddle.X - paddle.Width/2 - ball.Radius - 1;
   }
  
}