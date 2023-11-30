package gobases2

func GetSalary(hours int, category string) float32 {
	switch category {
	case "A":
		return calculateSalary(hours, 1000)
	case "B":
		return calculateSalary(hours, 1500) * 1.2
	case "C":
		return calculateSalary(hours, 3000) * 1.5
	default:
		return calculateSalary(hours, 1000)
	}
}

func calculateSalary(hours int, amount int) float32 {
	return float32((hours / 60) * amount)
}
