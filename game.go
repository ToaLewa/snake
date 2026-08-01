package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Game struct {
	squares  []MovingSquare
	snake    Snake
	food     Food
	gameOver bool
}

func NewGame() Game {
	return Game{
		squares: []MovingSquare{},
		snake: Snake{
			head:      gridRect(0, 60),
			tail:      []rl.Rectangle{gridRect(0, 40), gridRect(0, 20), gridRect(0, 0)},
			direction: East,
		},
		food: spawnFood(),
	}
}

func isOppositeDirection(currentDirection Direction, newDirection Direction) bool {
	isOpposite := false
	switch currentDirection {
	case North:
		if newDirection == South {
			isOpposite = true
		}
	case South:
		if newDirection == North {
			isOpposite = true
		}
	case East:
		if newDirection == West {
			isOpposite = true
		}
	case West:
		if newDirection == East {
			isOpposite = true
		}

	}

	return isOpposite
}

func UpdateGame(g *Game) {
	if g.gameOver {
		return
	}

	for i := range g.squares {
		moveAroundEdges(&g.squares[i])
	}

	moveByKeyboard(&g.snake)

	if float32(g.food.x) == g.snake.head.X && float32(g.food.y) == g.snake.head.Y {
		g.food = spawnFood()
		g.snake.tail = append(g.snake.tail, g.snake.tail[0])
	}

	for tailIndex := range g.snake.tail {
		if rl.CheckCollisionRecs(g.snake.tail[tailIndex], g.snake.head) {
			g.gameOver = true
			return
		}
	}

	updateSnakeTail(&g.snake)
	moveSnakeAroundEdges(&g.snake)

	/*for range 2 {
		g.squares = append(g.squares, spawnMovingSquare())
	}*/
}
