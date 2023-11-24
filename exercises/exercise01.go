package exercices

import (
	"strconv"
)

func IsMoreThan100(value string) (int, string) {
	var number, error = strconv.Atoi(value)
	if error == nil {
		if number < 100 {
			return number, "Is less than 100"
		} else {
			return number, "Is more than 100"
		}
	}
	return 0, "Number not valid"
}
