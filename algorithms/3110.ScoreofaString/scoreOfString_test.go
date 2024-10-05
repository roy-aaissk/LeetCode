package scoreofastring_test

import (
	scoreofastring "leet_code/algorithms/3110.ScoreofaString"
	"testing"
)

func Test_ScoreOfString(t *testing.T){
	tests := []struct{
		input string
		expected int
	}{
		{"hello", 13},
	}
	for _, test := range tests {
		result := scoreofastring.ScoreOfString(test.input)
		if result != test.expected {
			t.Errorf("ScoreOfString(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}
