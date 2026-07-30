package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func getInputDirection() (Direction, bool) {
	if rl.IsKeyPressed(rl.KeyUp) {
		return North, true
	} else if rl.IsKeyPressed(rl.KeyDown) {
		return South, true
	} else if rl.IsKeyPressed(rl.KeyLeft) {
		return West, true
	} else if rl.IsKeyPressed(rl.KeyRight) {
		return East, true
	} else {
		return "", false
	}
}
