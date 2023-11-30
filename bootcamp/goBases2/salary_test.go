package gobases2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSalary(t *testing.T) {
	t.Run("Validate A category", func(t *testing.T) {
		result := float32(192000)
		const (
			category = "A"
			minuts   = 11520
		)
		testResult := GetSalary(minuts, category)

		assert.Equal(t, result, testResult, "Result should be 192 000")
	})

	t.Run("Validate A category", func(t *testing.T) {
		result := float32(345600)
		const (
			category = "B"
			minuts   = 11520
		)
		testResult := GetSalary(minuts, category)

		assert.Equal(t, result, testResult, "Result should be 345 600")
	})

	t.Run("Validate A category", func(t *testing.T) {
		result := float32(864000)
		const (
			category = "C"
			minuts   = 11520
		)
		testResult := GetSalary(minuts, category)

		assert.Equal(t, result, testResult, "Result should be 864 600")
	})

	t.Run("Validate A category", func(t *testing.T) {
		result := float32(0)
		const (
			category = "j"
			minuts   = 11520
		)
		testResult := GetSalary(minuts, category)

		assert.Equal(t, result, testResult, "Result should be 0")
	})
}
