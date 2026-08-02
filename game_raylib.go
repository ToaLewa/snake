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

	renderFood(&g.food)
	drawSnake(&g.snake)

	if g.gameOver {
		drawGameOverScreen()
	}

	rl.EndDrawing()
}

func drawGameOverScreen() {
	overlayColor := rl.NewColor(0, 0, 0, 160)
	rl.DrawRectangle(0, 0, windowDetails.width, windowDetails.height, overlayColor)

	message := "GAME OVER"
	fontSize := int32(48)
	textWidth := rl.MeasureText(message, fontSize)
	x := (windowDetails.width - textWidth) / 2
	y := (windowDetails.height - fontSize) / 2

	rl.DrawText(message, x, y, fontSize, rl.White)
}

func renderFood(food *Food) {
	rl.DrawRectangle(food.x, food.y, gridIncrement, gridIncrement, rl.Red)
}

func renderMovingSquare(s *MovingSquare, color rl.Color) {
	rl.DrawRectangle(s.x, s.y, gridIncrement, gridIncrement, color)
}

func drawSnake(s *Snake) {
	for i := range s.tail {
		rl.DrawRectangleRec(s.tail[i], rl.Black)
	}

	rl.DrawRectangleRec(s.head, rl.Blue)
}
