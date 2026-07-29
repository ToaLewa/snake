package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type winStats struct {
	width  int32
	height int32
}

var windowDetails = winStats{width: 800, height: 400}

const fps = 20
const gridIncrement = 20

func onLeftBorder(x int32) bool {
	if x == 0 {
		return true
	} else {
		return false
	}
}

func onRightBorder(x int32) bool {
	if x == windowDetails.width-gridIncrement {
		return true
	} else {
		return false
	}
}

func onTopBorder(y int32) bool {
	if y == 0 {
		return true
	} else {
		return false
	}
}

func onBottomBorder(y int32) bool {
	if y == windowDetails.height-gridIncrement {
		return true
	} else {
		return false
	}
}

func main() {
	rl.InitWindow(windowDetails.width, windowDetails.height, "Snake")
	defer rl.CloseWindow()

	var y int32 = 0
	var x int32 = 0
	direction := "s"

	rl.SetTargetFPS(fps)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawRectangle(x, y, gridIncrement, gridIncrement, rl.Blue)
		rl.EndDrawing()

		if direction == "s" {
			if onBottomBorder(y) {
				direction = "e"
			} else {
				y += gridIncrement
			}

		} else if direction == "e" {
			if onRightBorder(x) {
				direction = "n"
			} else {
				x += gridIncrement
			}
		} else if direction == "n" {
			if onTopBorder(y) {
				direction = "w"
			} else {
				y -= gridIncrement
			}
		} else {
			if onLeftBorder(x) {
				direction = "s"
			} else {
				x -= gridIncrement
			}

		}

		fmt.Printf("%d and %d\n", y, windowDetails.height)

	}
}
