package Solutions

// Given a zero-based permutation nums (0-indexed), build an array ans of the same length where ans[i] = nums[nums[i]] for each 0 <= i < nums.length and return it.
// A zero-based permutation nums is an array of distinct integers from 0 to nums.length - 1 (inclusive).

func buildArray(nums []int) []int {
	var ans []int
	for _, num :=  range nums {
		ans = append(ans, nums[num])
	}
	return ans
}
