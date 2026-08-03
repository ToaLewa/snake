package main

import (
	"fmt"
	"os"
	"path/filepath"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	snake     Snake
	food      Food
	foodEaten int
	gameOver  bool
	borders   Borders
	shake     ScreenShake
}

type ScreenShake struct {
	frames    int
	magnitude float32
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
			growthRate: 5,
			direction:  East,
		},
		food:      spawnFood(),
		foodEaten: 0,
		borders: Borders{
			top:    rl.NewRectangle(0, -gridIncrement, float32(windowDetails.width), gridIncrement),
			bottom: rl.NewRectangle(0, float32(windowDetails.height), float32(windowDetails.width), gridIncrement),
			right:  rl.NewRectangle(float32(windowDetails.width), 0, gridIncrement, float32(windowDetails.height)),
			left:   rl.NewRectangle(-gridIncrement, 0, gridIncrement, float32(windowDetails.height)),
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
		ShakeScreen(g, 6, 8)
	}
}

func ShakeScreen(g *Game, frames int, magnitude float32) {
	g.shake.frames = frames
	g.shake.magnitude = magnitude
}

func updateScreenShake(g *Game) {
	if g.shake.frames > 0 {
		g.shake.frames--
	}
}

func savePath() (string, error) {
	configDir, err := os.UserConfigDir()

	if err != nil {
		return "", err
	}

	dir := filepath.Join(configDir, "power_snake")

	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}

	return filepath.Join(dir, "save.json"), nil
}

func Save(g *Game) error {
	path, err := savePath()

	if err != nil {
		return err
	}

	data := []byte(fmt.Sprintf(`{"high_score":%d}`, g.foodEaten))
	err = os.WriteFile(path, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func UpdateGame(g *Game) {
	updateScreenShake(g)
	checkTailCollision(g)

	if g.gameOver {
		if rl.IsKeyPressed(rl.KeyEnter) {
			Save(g)
			Restart(g)
		}

		return
	}

	pickDirectionByKeyboard(&g.snake)
	checkEat(g)

	updateSnakeTail(&g.snake)
	moveSnake(&g.snake)
	checkBorderCollision(g)
}

func checkTailCollision(g *Game) {
	for tailIndex := range g.snake.tail {
		if !g.gameOver && rl.CheckCollisionRecs(g.snake.tail[tailIndex], g.snake.head) {
			g.gameOver = true
			ShakeScreen(g, 6, 8)
		}
	}
}

func checkEat(g *Game) {
	if float32(g.food.x) == g.snake.head.X && float32(g.food.y) == g.snake.head.Y {
		g.food = spawnFood()
		for _ = range g.snake.growthRate {
			g.snake.tail = append(g.snake.tail, g.snake.tail[0])
		}
		g.foodEaten++
		ShakeScreen(g, 3, 1)
	}
}
