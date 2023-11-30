package gobases1

func CalculateGrade(grades ...int) int {
	if len(grades) == 0 {
		return 0
	}
	var resultado int
	for _, value := range grades {
		resultado += value
	}
	return resultado / len(grades)
}
