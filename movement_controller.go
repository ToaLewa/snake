package main

func moveByKeyboard(moveable *MovingSquare) {
	inputDirection, keyPressed := getInputDirection()
	if keyPressed {
		if !isOppositeDirection(moveable.direction, inputDirection) {
			moveable.direction = inputDirection
		}
	}
}
