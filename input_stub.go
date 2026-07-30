//go:build !raylib

package main

func getInputDirection() (Direction, bool) {
	return "", false
}
