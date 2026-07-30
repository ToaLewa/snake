package main

import (
	"math/rand"
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
	}
}

func main() {
	game := NewGame()

	SetupWindow()
	defer CloseWindow()

	for !WindowShouldClose() {
		UpdateGame(&game)
		DrawGame(&game)
	}
}
