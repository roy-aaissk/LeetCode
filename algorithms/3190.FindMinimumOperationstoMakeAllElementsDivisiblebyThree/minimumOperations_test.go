package findminimumoperationstomakeallelementsdivisiblebythree_test

import (
	findminimumoperationstomakeallelementsdivisiblebythree "leet_code/algorithms/3190.FindMinimumOperationstoMakeAllElementsDivisiblebyThree"
	"testing"
)

func Test_MinimumOperations(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{1, 2, 3, 4}, 3},
	}
	for _, test := range tests {
		result := findminimumoperationstomakeallelementsdivisiblebythree.MinimumOperations(test.input)
		if result != test.expected {
			t.Errorf("MinimumOperations(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}
