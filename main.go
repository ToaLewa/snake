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

func movingSquareRect(s *MovingSquare) rl.Rectangle {
	return rl.NewRectangle(float32(s.x), float32(s.y), gridIncrement, gridIncrement)
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
