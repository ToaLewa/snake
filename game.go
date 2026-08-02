package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Game struct {
	snake     Snake
	food      Food
	foodEaten int
	gameOver  bool
	borders   Borders
}

type Borders struct {
	top    rl.Rectangle
	bottom rl.Rectangle
	right  rl.Rectangle
	left   rl.Rectangle
}

func NewGame() Game {
	return Game{
		snake: Snake{
			head: gridRect(gridIncrement*5, gridIncrement),
			tail: []rl.Rectangle{
				gridRect(gridIncrement*4, gridIncrement),
				gridRect(gridIncrement*3, gridIncrement),
				gridRect(gridIncrement*2, gridIncrement),
				gridRect(gridIncrement*1, gridIncrement),
			},
			direction: East,
		},
		food:      spawnFood(),
		foodEaten: 0,
		borders: Borders{
			top:    rl.NewRectangle(0, 0, float32(windowDetails.width), gridIncrement),
			bottom: rl.NewRectangle(0, float32(windowDetails.height-gridIncrement), float32(windowDetails.width), gridIncrement),
			right:  rl.NewRectangle(float32(windowDetails.width-gridIncrement), 0, gridIncrement, float32(windowDetails.height)),
			left:   rl.NewRectangle(0, 0, gridIncrement, float32(windowDetails.height)),
		},
	}
}

func Restart(g *Game) {
	g.gameOver = false
	*g = NewGame()
}

func checkBorderCollision(g *Game) {
	if rl.CheckCollisionRecs(g.snake.head, g.borders.top) ||
		rl.CheckCollisionRecs(g.snake.head, g.borders.bottom) ||
		rl.CheckCollisionRecs(g.snake.head, g.borders.right) ||
		rl.CheckCollisionRecs(g.snake.head, g.borders.left) {
		g.gameOver = true
	}
}

func UpdateGame(g *Game) {
	checkTailCollision(g)
	checkBorderCollision(g)

	if g.gameOver {
		if rl.IsKeyPressed(rl.KeyEnter) {
			Restart(g)
		}

		return
	}

	pickDirectionByKeyboard(&g.snake)
	checkEat(g)

	updateSnakeTail(&g.snake)
	moveSnake(&g.snake)
}

func checkTailCollision(g *Game) {
	for tailIndex := range g.snake.tail {
		if rl.CheckCollisionRecs(g.snake.tail[tailIndex], g.snake.head) {
			g.gameOver = true
		}
	}
}

func checkEat(g *Game) {
	if float32(g.food.x) == g.snake.head.X && float32(g.food.y) == g.snake.head.Y {
		g.food = spawnFood()
		g.snake.tail = append(g.snake.tail, g.snake.tail[0])
	}
}
