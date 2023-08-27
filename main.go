package main

import (
	"github.com/NickDeChip/learning_supervised_learning/brain"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(600, 600, "Circles")
	rl.SetTargetFPS(30)

	var perceptron = brain.NewPerceptron()

	var points = make([]brain.Point, 200)
	for i := range points {
		points[i] = brain.NewPoint()
	}

	for !rl.WindowShouldClose() {

		if rl.IsKeyPressed(rl.KeyN) {
			for i := range points {
				perceptron.Train(points[i].GetLocationValue(), points[i].Value)
			}
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.Gray)
		for i := range points {
			points[i].Point(perceptron.Guess(points[i].GetLocationValue()))
		}
		rl.DrawLine(0, 600, 600, 0, rl.Black)

		rl.EndDrawing()
	}
	rl.CloseWindow()
}
