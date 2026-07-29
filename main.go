package main

/*
#cgo pkg-config: raylib
#include <raylib.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"

func main() {
	title := C.CString("Snake")
	defer C.free(unsafe.Pointer(title))

	C.InitWindow(800, 450, title)
	defer C.CloseWindow()

	C.SetTargetFPS(60)

	for !bool(C.WindowShouldClose()) {
		C.BeginDrawing()
		C.ClearBackground(C.RAYWHITE)
		C.DrawRectangle(0, 0, 20, 20, C.BLACK)
		// C.DrawRectangle(0, 20, 20, 20, C.BLACK)
		C.EndDrawing()
	}
}
