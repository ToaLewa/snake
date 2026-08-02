package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Game struct {
	snake     Snake
	food      Food
	foodEaten int
	gameOver  bool
}

func NewGame() Game {
	return Game{
		snake: Snake{
			head: gridRect(gridIncrement*5, 0),
			tail: []rl.Rectangle{
				gridRect(gridIncrement*4, 0),
				gridRect(gridIncrement*3, 0),
				gridRect(gridIncrement*2, 0),
				gridRect(gridIncrement*1, 0),
			},
			direction: East,
		},
		food:      spawnFood(),
		foodEaten: 0,
	}
}

func Restart(g *Game) {
	g.gameOver = false
	*g = NewGame()
}

func UpdateGame(g *Game) {
	if g.gameOver {
		if rl.IsKeyPressed(rl.KeyEnter) {
			Restart(g)
		}

		return
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
