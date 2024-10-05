package scoreofastring

import "math"

func ScoreOfString(s string) int {
	totalScore := 0
	for i := 0; i < len(s)-1; i++ {
		diff := int(s[i]) - int(s[i+1])
		totalScore += int(math.Abs(float64(diff)))
	}
	return totalScore
}
