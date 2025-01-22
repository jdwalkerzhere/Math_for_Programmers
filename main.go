package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(800, 450, "Snek")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)
	pos := rl.Vector3{X: 0, Y: 0, Z: 0}
	tar := rl.Vector3{X: 0, Y: 0, Z: -1}
	up := rl.Vector3{X: 0, Y: 1, Z: 0}
	camera := rl.NewCamera3D(pos, tar, up, 100, rl.CameraPerspective)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("Ok, now we've got a window folks!", 190, 200, 20, rl.LightGray)

		vect1 := rl.Vector2{X: 0, Y: 0}
		vect2 := rl.Vector2{X: 50, Y: 50}

		tdVect1 := rl.Vector3{X: 0, Y: 0, Z: 0}
		tdVect2 := rl.Vector3{X: 5, Y: 5, Z: 5}

		if rl.IsKeyDown(rl.KeyA) {
			rl.DrawLineV(vect1, vect2, rl.Black)
		}
		if rl.IsKeyDown(rl.KeyB) {
			rl.BeginMode3D(camera)
			rl.DrawCubeWiresV(tdVect1, tdVect2, rl.Black)
			rl.EndMode3D()
		}
		rl.EndDrawing()
	}
}
