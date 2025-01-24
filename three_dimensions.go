package main

import (
	"fmt"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Meteor struct {
	pos        rl.Vector3
	rad        float32
	color      rl.Color
	trajectory rl.Vector3
	velocity   int32
}

func main() {
	rl.InitWindow(800, 800, "Snek")
	defer rl.CloseWindow()

	threeDimensionalCamera := rl.Camera3D{
		Position:   rl.Vector3{X: 0, Y: 5, Z: -10},
		Target:     rl.Vector3{X: 0, Y: 0, Z: 400},
		Up:         rl.Vector3{X: 0, Y: 1, Z: 0},
		Fovy:       90,
		Projection: rl.CameraPerspective,
	}

	meteors := make([]*Meteor, 10)
	for i := range 10 {
		meteors[i] = &Meteor{
			pos: rl.Vector3{X: float32(rl.GetRandomValue(-400, 400)),
				Y: float32(rl.GetRandomValue(350, 400)),
				Z: float32(rl.GetRandomValue(-400, 400))},
			rad:   float32(rl.GetRandomValue(5, 50)),
			color: rl.Red,
			trajectory: rl.Vector3{
				X: float32(rl.GetRandomValue(-1, 1)),
				Y: float32(rl.GetRandomValue(-2, -1)),
				Z: float32(rl.GetRandomValue(-1, 1))},
			velocity: rl.GetRandomValue(1, 3),
		}
	}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.White)
		rl.UpdateCamera(&threeDimensionalCamera, rl.CameraFirstPerson)
		rl.BeginMode3D(threeDimensionalCamera)
		// rl.DrawGrid(400, 1.0)
		rl.DrawPlane(rl.Vector3{X: 0, Y: 0, Z: 0}, rl.Vector2{X: 800, Y: 800}, rl.Green)
		rl.DrawLine3D(
			rl.Vector3{X: 0, Y: -400, Z: 0},
			rl.Vector3{X: 0, Y: 400, Z: 0},
			rl.Red)
		rl.DrawCube(rl.Vector3{X: 0, Y: 0, Z: 0}, 1, 1, 1, rl.Blue)

		if rl.IsKeyDown(rl.KeyB) {
			for _, m := range meteors {
				rl.DrawSphere(m.pos, m.rad, m.color)
				for range m.velocity {
					m.pos = rl.Vector3Add(m.pos, m.trajectory)
				}
				fmt.Println(m)
			}
			time.Sleep(10000000)
		}
		rl.EndMode3D()
		rl.EndDrawing()
	}
}
