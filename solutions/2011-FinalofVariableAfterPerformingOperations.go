package Solutions

func finalValueAfterOperations(operations []string) int {
	var keys []int
	var result int
	for k := range operations {
		keys = append(keys, k)
	}
	for _, key := range keys {
		if operations[key] == "--X" || operations[key] == "X--" {
			result--
		} else if operations[key] == "X++" || operations[key] == "++X" {
			result++
		}
	}
	return result
}
