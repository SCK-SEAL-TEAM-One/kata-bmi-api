package bmi

import "testing"

func Test_CalculateBMI_Input_Weight_83_Height_170_Should_Be_28dot72_And_Fat(t *testing.T) {
	weight := 83.00
	height := 170.00
	expected := 28.72

	actual := CalculateBMI(weight, height)

	if expected != actual {
		t.Errorf("expected is %.2f but got %.2f", expected, actual)
	}
}
