package buildarrayfrompermutation

func BuildArray(nums []int) []int {
	result := make([]int, len(nums))
	for i, num := range nums {
			result[i] = nums[num]
	}
	return result
}
