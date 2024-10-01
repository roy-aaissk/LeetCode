package palindromenumber_test

import (
	testTarget "leet_code/algorithms/9.PalindromeNumber"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input    int
		expected bool
	}{
		{121, true},
		{-121, false},
		{10, false},
		{1, true},
		{12321, true},
		{0, true},
	}

	for _, test := range tests {
		result := testTarget.IsPalindrome(test.input)
		if result != test.expected {
			t.Errorf("isPalindrome(%d) = %v; want %v", test.input, result, test.expected)
		}
	}
}
