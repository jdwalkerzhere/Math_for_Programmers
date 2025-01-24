package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 800, "Snek")
	defer rl.CloseWindow()

	twoDimensionalCamera := rl.Camera2D{
		Target:   rl.Vector2{X: 0, Y: 0},
		Offset:   rl.Vector2{X: 400, Y: 400},
		Rotation: 0,
		Zoom:     1.0,
	}

	threeDimensionalCamera := rl.Camera3D{
		Position:   rl.Vector3{X: 0, Y: 5, Z: -10},
		Target:     rl.Vector3{X: 0, Y: 0, Z: 400},
		Up:         rl.Vector3{X: 0, Y: 1, Z: 0},
		Fovy:       90,
		Projection: rl.CameraPerspective,
	}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.White)
		if rl.IsKeyDown(rl.KeyK) {
			rl.BeginMode2D(twoDimensionalCamera)
			rl.DrawLineV(rl.Vector2{X: -400, Y: 0}, rl.Vector2{X: 400, Y: 0}, rl.Black)
			rl.DrawLineV(rl.Vector2{X: 0, Y: -400}, rl.Vector2{X: 0, Y: 400}, rl.Black)
			rl.EndMode2D()
		}
		if rl.IsKeyDown(rl.KeyB) {
			rl.UpdateCamera(&threeDimensionalCamera, rl.CameraFirstPerson)
			rl.BeginMode3D(threeDimensionalCamera)
			rl.DrawGrid(400, 1.0)
			rl.DrawLine3D(
				rl.Vector3{X: 0, Y: -400, Z: 0},
				rl.Vector3{X: 0, Y: 400, Z: 0},
				rl.Red)
			rl.DrawCube(rl.Vector3{X: 0, Y: 0, Z: 0}, 1, 1, 1, rl.Blue)
			rl.EndMode3D()
		}
		rl.EndDrawing()
	}
}
