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

func UpdateGame(g *Game) {
	for i := range g.squares {
		moveAroundEdges(&g.squares[i])
	}

	inputDirection, keyPressed := getInputDirection()
	if keyPressed {
		g.snake.head.direction = inputDirection
	}

	updateSnakeTail(&g.snake)
	moveAroundEdges(&g.snake.head)

	if g.food.x == g.snake.head.x && g.food.y == g.snake.head.y {
		g.food = spawnFood()
	}

	/*for range 2 {
		g.squares = append(g.squares, spawnMovingSquare())
	}*/
}
