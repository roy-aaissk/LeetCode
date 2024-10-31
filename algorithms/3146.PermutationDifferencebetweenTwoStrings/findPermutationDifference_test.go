package findPermutationDifference_test

import (
	testTarget "leet_code/algorithms/3146.PermutationDifferencebetweenTwoStrings"
	"testing"
)

func Test_findPermutationDifference(t *testing.T) {
	tests := []struct {
		s        string
		t        string
		expected int
	}{
		{s: "abc", t: "bac", expected: 2},
	}
	for _, test := range tests {
		result := testTarget.FindPermutationDifference(test.s, test.t)
		if result != test.expected {
			t.Errorf("FindPermutationDifference(%v) = %v; want %v", test.s, result, test.expected)
		}
	}
}
