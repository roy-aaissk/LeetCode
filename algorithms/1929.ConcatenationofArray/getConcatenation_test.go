package concatenationofarray_test

import (
	concatenationofarray "leet_code/algorithms/1929.ConcatenationofArray"
	"reflect"
	"testing"
)

func Test_GetConcatenation(t *testing.T) {
	tests := []struct {
		nums     []int
		expected []int
	}{
		{[]int{1,2,1}, []int{1,2,1,1,2,1}},
		{[]int{1,3,2,1}, []int{1,3,2,1,1,3,2,1}},
	}
	for _, test := range tests {
		result := concatenationofarray.GetConcatenation(test.nums)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("GetConcatenation(%v) = %v; want %v", test.nums, result, test.expected)
		}
	}
}
