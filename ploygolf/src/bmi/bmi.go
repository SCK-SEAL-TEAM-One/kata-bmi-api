package bmi

import "math"

func CalculateBmi(height, weight float64) float64 {
	heightMeter := height / 100
	bmi := weight / (heightMeter * heightMeter)
	bmi = math.Round(bmi*100) / 100
	return bmi
}
