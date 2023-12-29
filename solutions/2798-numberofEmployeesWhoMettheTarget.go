package Solutions

func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
	var result int
	for _, hour := range hours {
		if hour >= target {
			result += 1
		}
	}
	return result
}
