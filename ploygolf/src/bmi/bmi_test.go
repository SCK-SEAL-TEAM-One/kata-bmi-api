package bmi

import "testing"

func Test_CalculateBmi_Input_170_83_Should_Be_28_Dot_72(t *testing.T) {
	height := 170.00
	weight := 83.00
	expected := 28.72

	actualResult := CalculateBmi(height, weight)

	if expected != actualResult {
		t.Errorf("Expect %f but got %f", expected, actualResult)
	}
}
