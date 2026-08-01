package main

func moveByKeyboard(s *Snake) {
	inputDirection, keyPressed := getInputDirection()
	if keyPressed {
		if !isOppositeDirection(s.direction, inputDirection) {
			s.direction = inputDirection
		}
	}
}
