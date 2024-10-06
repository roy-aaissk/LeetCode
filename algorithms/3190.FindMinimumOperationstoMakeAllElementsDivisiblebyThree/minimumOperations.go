package findminimumoperationstomakeallelementsdivisiblebythree

func MinimumOperations(nums []int) int {
	result := 0
	for _, num := range nums {
		if num%3 == 0 {
			continue
		} else if num%3 == 1 {
			result++
		} else if num%3 == 2 {
			result++
		}
	}
	return result
}
