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

type Food struct {
	x     int32
	y     int32
	color rl.Color
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
	head      rl.Rectangle
	tail      []rl.Rectangle
	direction Direction
}

var windowDetails = WinStats{width: 800, height: 400}

const fps = 20
const gridIncrement = 20

func movingSquareRect(s *MovingSquare) rl.Rectangle {
	return rl.NewRectangle(float32(s.x), float32(s.y), gridIncrement, gridIncrement)
}

func gridRect(x int32, y int32) rl.Rectangle {
	return rl.NewRectangle(float32(x), float32(y), gridIncrement, gridIncrement)
}

func leftBorderRect() rl.Rectangle {
	return rl.NewRectangle(0, 0, gridIncrement, float32(windowDetails.height))
}

func rightBorderRect() rl.Rectangle {
	return rl.NewRectangle(float32(windowDetails.width-gridIncrement), 0, gridIncrement, float32(windowDetails.height))
}

func topBorderRect() rl.Rectangle {
	return rl.NewRectangle(0, 0, float32(windowDetails.width), gridIncrement)
}

func bottomBorderRect() rl.Rectangle {
	return rl.NewRectangle(0, float32(windowDetails.height-gridIncrement), float32(windowDetails.width), gridIncrement)
}

func moveAroundEdges(s *MovingSquare) {
	squareRect := movingSquareRect(s)

	if s.direction == South {
		if rl.CheckCollisionRecs(squareRect, bottomBorderRect()) {
			s.direction = East
		} else {
			s.y += gridIncrement
		}

	} else if s.direction == East {
		if rl.CheckCollisionRecs(squareRect, rightBorderRect()) {
			s.direction = North
		} else {
			s.x += gridIncrement
		}
	} else if s.direction == North {
		if rl.CheckCollisionRecs(squareRect, topBorderRect()) {
			s.direction = West
		} else {
			s.y -= gridIncrement
		}
	} else {
		if rl.CheckCollisionRecs(squareRect, leftBorderRect()) {
			s.direction = South
		} else {
			s.x -= gridIncrement
		}

	}
}

func moveSnakeAroundEdges(s *Snake) {
	if s.direction == South {
		if rl.CheckCollisionRecs(s.head, bottomBorderRect()) {
			s.direction = East
		} else {
			s.head.Y += gridIncrement
		}

	} else if s.direction == East {
		if rl.CheckCollisionRecs(s.head, rightBorderRect()) {
			s.direction = North
		} else {
			s.head.X += gridIncrement
		}
	} else if s.direction == North {
		if rl.CheckCollisionRecs(s.head, topBorderRect()) {
			s.direction = West
		} else {
			s.head.Y -= gridIncrement
		}
	} else {
		if rl.CheckCollisionRecs(s.head, leftBorderRect()) {
			s.direction = South
		} else {
			s.head.X -= gridIncrement
		}

	}
}

func spawnMovingSquare() MovingSquare {
	directions := []Direction{South, East, North, West}

	coord := getRandomCoordinate()

	return MovingSquare{
		x:         coord.x,
		y:         coord.y,
		direction: directions[rand.Intn(len(directions))],
	}
}

func spawnFood() Food {
	coord := getRandomCoordinate()
	return Food{
		x: coord.x,
		y: coord.y,
	}
}

type coord struct {
	x int32
	y int32
}

func getRandomCoordinate() coord {
	return coord{
		x: rand.Int31n(windowDetails.width/gridIncrement) * gridIncrement,
		y: rand.Int31n(windowDetails.height/gridIncrement) * gridIncrement,
	}
}

func updateSnakeTail(s *Snake) {
	prevX, prevY := s.head.X, s.head.Y

	for i := range s.tail {
		oldX, oldY := s.tail[i].X, s.tail[i].Y

		s.tail[i].X = prevX
		s.tail[i].Y = prevY
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
