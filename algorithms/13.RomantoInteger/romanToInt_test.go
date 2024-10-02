package romantointeger_test

import (
	testTarget "leet_code/algorithms/13.RomantoInteger"
	"testing"
)

func Test_RomanToInt(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"III", 3},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
	}
	for _, test := range tests {
		result := testTarget.RomanToInt(test.input)
		if test.expected != result {
			t.Errorf("RomanToInt(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}
