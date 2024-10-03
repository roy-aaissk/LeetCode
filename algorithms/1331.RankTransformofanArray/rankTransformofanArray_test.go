package ranktransformofanarray_test

import (
	ranktransformofanarray "leet_code/algorithms/1331.RankTransformofanArray"
	"reflect"
	"testing"
)

func Test_ArrayRankTransform(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{40, 10, 20, 30}, []int{4, 1, 2, 3}},
		{[]int{100, 100, 100}, []int{1, 1, 1}},
		{[]int{37, 12, 28, 9, 100, 56, 80, 5, 12}, []int{5, 3, 4, 2, 8, 6, 7, 1, 3}},
	}
	for _, test := range tests {
		results := ranktransformofanarray.ArrayRankTransform(test.input)
		if !reflect.DeepEqual(results, test.expected) {
			t.Errorf("ArrayRankTransform(%v) = %v; want %v", test.input, results, test.expected)
		}
	}
}
