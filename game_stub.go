//go:build !raylib

package main

func SetupGame() Game {
	return NewGame()
}

func CloseGame() {
}

func WindowShouldClose() bool {
	return true
}

func DrawGame(g *Game) {
}
