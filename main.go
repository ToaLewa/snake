package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(800, 400, "Snake")
	defer rl.CloseWindow()

	var y int32 = 0

	rl.SetTargetFPS(10)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawRectangle(0, y, 20, 20, rl.Black)
		// rl.DrawRectangle(0, 20, 20, 20, rl.Black)
		rl.EndDrawing()

		y += 20
	}
}
