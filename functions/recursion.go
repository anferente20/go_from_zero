package functions

func Summatory(limit int) int {
	if limit == 1 {
		return limit
	} else {
		return limit + Summatory(limit-1)
	}
}
