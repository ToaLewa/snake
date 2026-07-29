package main

import (
	// "fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type winStats struct {
	width  int32
	height int32
}

var windowDetails = winStats{width: 800, height: 400}

const gridIncrement = 20

func onBorder(x int, y int) bool {
	return true
}

func main() {
	rl.InitWindow(windowDetails.width, windowDetails.height, "Snake")
	defer rl.CloseWindow()

	var y int32 = 0
	var x int32 = 0
	// direction := "s"

	rl.SetTargetFPS(10)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawRectangle(x, y, gridIncrement, gridIncrement, rl.Black)
		rl.EndDrawing()

		bottom := y == windowDetails.height-gridIncrement

		if !bottom {
			y += gridIncrement
		} else {
			x += gridIncrement
		}

		// fmt.Printf("%d\n", y)

	}
}
