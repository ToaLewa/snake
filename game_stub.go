//go:build !raylib

package main

func SetupWindow() {
}

func CloseWindow() {
}

func WindowShouldClose() bool {
	return true
}

func DrawGame(g *Game) {
}
