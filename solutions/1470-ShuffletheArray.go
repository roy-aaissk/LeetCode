package Solutions

// Given the array nums consisting of 2n elements in the form [x1,x2,...,xn,y1,y2,...,yn].

// Return the array in the form [x1,y1,x2,y2,...,xn,yn].

func shuffle(nums []int, n int) []int {
	sliceFirst := nums[0:n]
	sliceLatter := nums[n:]
	var result []int
	for i := 0; i < n; i++ {
		result = append(result, sliceFirst[i], sliceLatter[i])
	}
	return result
}
