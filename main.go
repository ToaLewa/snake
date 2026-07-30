package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Direction string

const (
	South Direction = "s"
	East  Direction = "e"
	North Direction = "n"
	West  Direction = "w"
)

type FollowerSquare struct {
	x int32
	y int32
}

type MovingSquare struct {
	x         int32
	y         int32
	direction Direction
}

type WinStats struct {
	width  int32
	height int32
}

type Snake struct {
	head MovingSquare
	tail []FollowerSquare
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

func spawnMovingSquare() MovingSquare {
	directions := []Direction{South, East, North, West}

	return MovingSquare{
		x:         rand.Int31n(windowDetails.width/gridIncrement) * gridIncrement,
		y:         rand.Int31n(windowDetails.height/gridIncrement) * gridIncrement,
		direction: directions[rand.Intn(len(directions))],
	}
}

func updateSnakeTail(s *Snake) {
	prevX, prevY := s.head.x, s.head.y

	for i := range s.tail {
		oldX, oldY := s.tail[i].x, s.tail[i].y

		s.tail[i].x = prevX
		s.tail[i].y = prevY

		prevX, prevY = oldX, oldY

		rl.DrawRectangle(s.tail[i].x, s.tail[i].y, gridIncrement, gridIncrement, rl.Blue)
	}
}

func main() {
	squares := []MovingSquare{}

	snake := Snake{
		head: MovingSquare{x: 0, y: 60, direction: East},
		tail: []FollowerSquare{{x: 0, y: 40}, {x: 0, y: 20}, {x: 0, y: 0}},
	}

	rl.InitWindow(windowDetails.width, windowDetails.height, "Snake")
	defer rl.CloseWindow()
	rl.SetTargetFPS(fps)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.EndDrawing()

		for i := range squares {
			var s *MovingSquare = &squares[i]
			moveAroundEdges(s)
			renderMovingSquare(s)
		}

		inputDirection, keyPressed := getInputDirection()

		if keyPressed {
			snake.head.direction = inputDirection
		}

		updateSnakeTail(&snake)
		var s *MovingSquare = &snake.head
		moveAroundEdges(s)
		renderMovingSquare(s)

		for i := 0; i < 2; i++ {
			squares = append(squares, spawnMovingSquare())

		}
		// fmt.Printf("%d and %d\n", sy, windowDetails.height)

	}
}
