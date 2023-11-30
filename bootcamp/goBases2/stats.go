package gobases2

import "fmt"

const (
	minimum = "minimum"
	maximum = "maximum"
	average = "average"
)

func getMinimum(values ...float32) (float32, error) {
	if len(values) == 0 {
		return 0, fmt.Errorf("Grades are empty.")
	}
	minimum := values[0]
	for _, value := range values {
		if value < minimum {
			minimum = value
		}
	}
	return minimum, nil
}

func getMaximum(values ...float32) (float32, error) {
	if len(values) == 0 {
		return 0, fmt.Errorf("Grades are empty.")
	}
	maximum := values[0]
	for _, value := range values {
		if value > maximum {
			maximum = value
		}
	}
	return maximum, nil
}

func getAverage(values ...float32) (float32, error) {
	if len(values) == 0 {
		return 0, fmt.Errorf("Grades are empty.")
	}
	summatory := float32(0.0)

	for _, value := range values {
		summatory += value
	}
	return summatory / float32(len(values)), nil
}

func GetFunction(operation string) func(values ...float32) (float32, error) {
	switch operation {
	case minimum:
		return getMinimum
	case maximum:
		return getMaximum
	case average:
		return getAverage
	default:
		return getAverage
	}
}
