package main

import rl "github.com/gen2brain/raylib-go/raylib"

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
		if rl.CheckCollisionRecs(s.head, bottomBorderRect()) {
			s.direction = East
		} else {
			s.head.Y += gridIncrement
		}

	} else if s.direction == East {
		if rl.CheckCollisionRecs(s.head, rightBorderRect()) {
			s.direction = North
		} else {
			s.head.X += gridIncrement
		}
	} else if s.direction == North {
		if rl.CheckCollisionRecs(s.head, topBorderRect()) {
			s.direction = West
		} else {
			s.head.Y -= gridIncrement
		}
	} else {
		if rl.CheckCollisionRecs(s.head, leftBorderRect()) {
			s.direction = South
		} else {
			s.head.X -= gridIncrement
		}

	}
}
