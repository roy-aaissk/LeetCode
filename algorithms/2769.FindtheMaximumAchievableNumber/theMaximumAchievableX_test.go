package findthemaximumachievablenumber_test

import (
	findthemaximumachievablenumber "leet_code/algorithms/2769.FindtheMaximumAchievableNumber"
	"testing"
)

func Test_TheMaximumAchievableX(t *testing.T) {
	tests := []struct {
		num      int
		t        int
		expected int
	}{
		{4, 1, 6},
		{3, 2, 7},
	}
	for _, test := range tests {
		result := findthemaximumachievablenumber.TheMaximumAchievableX(test.num, test.t)
		if test.expected != result {
			t.Errorf("TheMaximumAchievableX(%v) = %v; want %v", test.num, result, test.expected)
		}
	}
}
