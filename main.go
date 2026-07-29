package main

import (
	// "fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Direction string

const (
	South Direction = "s"
	East  Direction = "e"
	North Direction = "n"
	West  Direction = "w"
)

type MovingSquare struct {
	x         int32
	y         int32
	direction Direction
}

type WinStats struct {
	width  int32
	height int32
}

var windowDetails = WinStats{width: 800, height: 400}

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

func moveAroundEdges(s *MovingSquare) {
	if s.direction == South {
		if onBottomBorder(s.y) {
			s.direction = East
		} else {
			s.y += gridIncrement
		}

	} else if s.direction == East {
		if onRightBorder(s.x) {
			s.direction = North
		} else {
			s.x += gridIncrement
		}
	} else if s.direction == North {
		if onTopBorder(s.y) {
			s.direction = West
		} else {
			s.y -= gridIncrement
		}
	} else {
		if onLeftBorder(s.x) {
			s.direction = South
		} else {
			s.x -= gridIncrement
		}

	}
}

func renderMovingSquare(s *MovingSquare) {
	rl.DrawRectangle(s.x, s.y, gridIncrement, gridIncrement, rl.Blue)
}

func updateMovingSquare(s *MovingSquare) {
	moveAroundEdges(s)
	renderMovingSquare(s)
}

func main() {
	rl.InitWindow(windowDetails.width, windowDetails.height, "Snake")
	defer rl.CloseWindow()

	square1 := MovingSquare{
		x:         0,
		y:         0,
		direction: South,
	}

	square2 := MovingSquare{
		x:         0,
		y:         20,
		direction: South,
	}

	rl.SetTargetFPS(fps)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.EndDrawing()

		updateMovingSquare(&square1)
		updateMovingSquare(&square2)
		// fmt.Printf("%d and %d\n", sy, windowDetails.height)

	}
}
