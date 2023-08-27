package brain

import rl "github.com/gen2brain/raylib-go/raylib"

type Point struct {
	X     float32
	Y     float32
	Value int32
}

func NewPoint() Point {
	point := Point{}
	point.X = getRandFloatVal(-1, 1)
	point.Y = getRandFloatVal(-1, 1)

	if point.X > point.Y {
		point.Value = 1
	} else {
		point.Value = -1
	}

	return point
}

func (point *Point) GetPixelX() int32 {
	return int32(map_value(float32(point.X), -1, 1, 0, 600))
}

func (point *Point) GetPixelY() int32 {
	return int32(map_value(float32(point.Y), -1, 1, 600, 0))
}

func (point *Point) Point(guess float32) {
	rl.DrawCircle(point.GetPixelX(), point.GetPixelY(), 18, rl.Black)
	if guess == float32(point.Value) {
		rl.DrawCircle(point.GetPixelX(), point.GetPixelY(), 14, rl.Green)
	} else {
		rl.DrawCircle(point.GetPixelX(), point.GetPixelY(), 14, rl.Red)
	}
}

func map_value(n, start1, stop1, start2, stop2 float32) float32 {
	return ((n-start1)/(stop1-start1))*(stop2-start2) + start2
}

func (point *Point) GetLocationValue() []float32 {
	input := make([]float32, 2)
	input[0] = float32(point.X)
	input[1] = float32(point.Y)
	return input
}
