package gobases2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFunction(t *testing.T) {
	t.Run("Validate minimum value", func(t *testing.T) {

		result := float32(4)
		const operation = "minimum"
		grades := []float32{90, 80, 100, 4, 12, 33}

		minFunc := GetFunction(operation)
		minVal, err := minFunc(grades...)
		if err != nil {
			t.Errorf(err.Error())
		}

		assert.Equal(t, result, minVal, "Result should be 4")
	})

	t.Run("Validate maximum value", func(t *testing.T) {

		result := float32(100)
		const operation = "maximum"
		grades := []float32{90, 80, 100, 4, 12, 33}

		maxFunc := GetFunction(operation)
		maxVal, err := maxFunc(grades...)
		if err != nil {
			t.Errorf(err.Error())
		}

		assert.Equal(t, result, maxVal, "Result should be 4")
	})

	t.Run("Validate average value", func(t *testing.T) {

		result := float32(53.1666667)
		const operation = "average"
		grades := []float32{90, 80, 100, 4, 12, 33}

		aveFunc := GetFunction(operation)
		aveVal, err := aveFunc(grades...)
		if err != nil {
			t.Errorf(err.Error())
		}

		assert.Equal(t, result, aveVal, "Result should be 4")
	})
}
