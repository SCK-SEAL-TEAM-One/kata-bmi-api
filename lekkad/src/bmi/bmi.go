package bmi

import "math"

type Bmi struct {
	Bmi    float64
	Status string
}

const MATE = 100

func CalculateBMI(weight, height float64) float64 {
	heightMate := height / MATE
	bmi := weight / (heightMate * heightMate)

	return math.Round(bmi*100) / 100
}
