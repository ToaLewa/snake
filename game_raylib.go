package main

import rl "github.com/gen2brain/raylib-go/raylib"

func SetupWindow() {
	rl.InitWindow(windowDetails.width, windowDetails.height, "Snake")
	rl.SetTargetFPS(fps)
}

func CloseWindow() {
	rl.CloseWindow()
}

func WindowShouldClose() bool {
	return rl.WindowShouldClose()
}

func DrawGame(g *Game) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	for i := range g.squares {
		renderMovingSquare(&g.squares[i], rl.Blue)
	}

	drawSnake(&g.snake)

	rl.EndDrawing()
}

func renderMovingSquare(s *MovingSquare, color rl.Color) {
	rl.DrawRectangle(s.x, s.y, gridIncrement, gridIncrement, color)
}

func drawSnake(s *Snake) {
	for i := range s.tail {
		rl.DrawRectangle(s.tail[i].x, s.tail[i].y, gridIncrement, gridIncrement, rl.Black)
	}

	renderMovingSquare(&s.head, rl.Black)
}
