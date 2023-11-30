package gobases1

import (
	"strconv"
	"testing"
)

// Validate returnb when len(grades) == 0
func TestCalculateGrade(t *testing.T) {
	t.Run("Validate when grades variable is empty", //Arrange
		func(t *testing.T) {
			//Arrange
			result := 0
			var grades []int

			testResult := CalculateGrade(grades...)
			if result != testResult {
				t.Errorf("Should respond 0. Waiting... Expected = %s, Obtained  = %s", strconv.Itoa(result), strconv.Itoa(testResult))
			}
		})

	t.Run("Validates average",
		func(t *testing.T) {
			//Arrange
			result := 5
			grades := []int{2, 4, 6, 8}

			testResult := CalculateGrade(grades...)
			if result != testResult {
				t.Errorf("Should respond 0. Waiting... Expected = %s, Obtained  = %s", strconv.Itoa(result), strconv.Itoa(testResult))
			}
		},
	)

}
