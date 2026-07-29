package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(800, 450, "Snake")
	defer rl.CloseWindow()

	rl.SetTargetFPS(1)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawRectangle(0, 0, 20, 20, rl.Black)
		// rl.DrawRectangle(0, 20, 20, 20, rl.Black)
		rl.EndDrawing()
	}
}
