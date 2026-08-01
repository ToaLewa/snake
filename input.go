package main

import rl "github.com/gen2brain/raylib-go/raylib"

func getInputDirection() (Direction, bool) {
	if rl.IsKeyPressed(rl.KeyUp) || rl.IsKeyPressed(rl.KeyW) {
		return North, true
	} else if rl.IsKeyPressed(rl.KeyDown) || rl.IsKeyPressed(rl.KeyS) {
		return South, true
	} else if rl.IsKeyPressed(rl.KeyLeft) || rl.IsKeyPressed(rl.KeyA) {
		return West, true
	} else if rl.IsKeyPressed(rl.KeyRight) || rl.IsKeyPressed(rl.KeyD) {
		return East, true
	} else {
		return "", false
	}
}
