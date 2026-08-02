package main

func isOppositeDirection(currentDirection Direction, newDirection Direction) bool {
	isOpposite := false
	switch currentDirection {
	case North:
		if newDirection == South {
			isOpposite = true
		}
	case South:
		if newDirection == North {
			isOpposite = true
		}
	case East:
		if newDirection == West {
			isOpposite = true
		}
	case West:
		if newDirection == East {
			isOpposite = true
		}

	}

	return isOpposite
}

func pickDirectionByKeyboard(s *Snake) {
	inputDirection, keyPressed := getInputDirection()
	if keyPressed {
		if !isOppositeDirection(s.direction, inputDirection) {
			s.direction = inputDirection
		}
	}
}

func moveSnake(s *Snake) {
	if s.direction == South {
		s.head.Y += gridIncrement
	} else if s.direction == East {
		s.head.X += gridIncrement
	} else if s.direction == North {
		s.head.Y -= gridIncrement
	} else {
		s.head.X -= gridIncrement
	}
}
