package brain

import (
	"math/rand"
)

const learningRate = 0.2

type Perceptron struct {
	Weights []float32
}

func NewPerceptron() Perceptron {
	brain := Perceptron{}
	brain.Weights = make([]float32, 2)
	for i := range brain.Weights {
		brain.Weights[i] = getRandFloatVal(-1, 1)
	}
	return brain
}

func (brain *Perceptron) Guess(inputs []float32) float32 {
	var sum float32
	for i := range brain.Weights {
		sum += inputs[i] * brain.Weights[i]
	}
	return process(sum)
}

func (brain *Perceptron) Train(inputs []float32, target int32) {
	guess := brain.Guess(inputs)
	error := target - int32(guess)

	for i := range brain.Weights {
		brain.Weights[i] += float32(error) * inputs[i] * learningRate
	}
}

func process(val float32) float32 {
	if val >= 0 {
		return 1
	} else {
		return -1
	}
}

func getRandFloatVal(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}
