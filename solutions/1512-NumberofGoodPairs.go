package Solutions

func numIdenticalPairs(nums []int) int {
	totalNums := len(nums)
	var result int
	for index, num := range nums {
			for i := index+1; i < totalNums; i++ {
					if num == nums[i] {
							result++
					}
			}
	}
	return result
}
