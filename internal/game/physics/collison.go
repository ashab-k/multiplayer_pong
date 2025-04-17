package physics

import (
	"ashab-k/github.com/internal/game/entity"
	"fmt"
)

func CheckPaddleBallCollision(ball *entity.Ball, paddle *entity.Paddle) bool {
    paddleLeft := paddle.X
    paddleRight := paddle.X + paddle.Width
    paddleTop := paddle.Y
    paddleBottom := paddle.Y + paddle.Height

    if ball.X+ball.Radius >= paddleLeft &&
        ball.X-ball.Radius <= paddleRight &&
        ball.Y+ball.Radius >= paddleTop &&
        ball.Y-ball.Radius <= paddleBottom {

        if paddle.X < 320/2 && ball.Dx < 0 {
            handlePaddleBallResponse(ball, paddle)
            return true
        } else if paddle.X > 320/2 && ball.Dx > 0 {
            handlePaddleBallResponse(ball, paddle)
            return true
        }
    }

    return false
}


func handlePaddleBallResponse(ball *entity.Ball, paddle *entity.Paddle) {
    fmt.Println("paddle ball collision")
    ball.Dx = -ball.Dx 

    paddleCenterY := paddle.Y + paddle.Height/2
    hitPosition := (ball.Y - paddleCenterY) / (paddle.Height / 2)
    ball.Dy = hitPosition * 2 

    if paddle.X < 160 {
        ball.X = paddle.X + paddle.Width + ball.Radius + 1
    } else {
        ball.X = paddle.X - ball.Radius - 1
    }
}
