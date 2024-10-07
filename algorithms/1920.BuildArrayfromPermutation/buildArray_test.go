package buildarrayfrompermutation_test

import (
	testTarget "leet_code/algorithms/1920.BuildArrayfromPermutation"
	"reflect"
	"testing"
)

func Test_BuildArray(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{0, 2, 1, 5, 3, 4}, []int{0, 1, 2, 4, 5, 3}},
		{[]int{5, 0, 1, 2, 3, 4}, []int{4, 5, 0, 1, 2, 3}},
	}
	for _, test := range tests {
		results := testTarget.BuildArray(test.input)
		if !reflect.DeepEqual(results, test.expected) {
			t.Errorf("BuildArray(%v) = %v; want %v", test.input, results, test.expected)
		}
	}
}
