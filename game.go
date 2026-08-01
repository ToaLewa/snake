package main

type Game struct {
	squares []MovingSquare
	snake   Snake
	food    Food
}

func NewGame() Game {
	return Game{
		squares: []MovingSquare{},
		snake: Snake{
			head: MovingSquare{x: 0, y: 60, direction: East},
			tail: []FollowerSquare{{x: 0, y: 40}, {x: 0, y: 20}, {x: 0, y: 0}},
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
	for i := range g.squares {
		moveAroundEdges(&g.squares[i])
	}

	inputDirection, keyPressed := getInputDirection()
	if keyPressed {
		if !isOppositeDirection(g.snake.head.direction, inputDirection) {
			g.snake.head.direction = inputDirection
		}
	}

	if g.food.x == g.snake.head.x && g.food.y == g.snake.head.y {
		g.food = spawnFood()
		g.snake.tail = append(g.snake.tail, FollowerSquare{x: g.snake.tail[0].x, y: g.snake.tail[0].y})
	}

	updateSnakeTail(&g.snake)
	moveAroundEdges(&g.snake.head)

	/*for range 2 {
		g.squares = append(g.squares, spawnMovingSquare())
	}*/
}
