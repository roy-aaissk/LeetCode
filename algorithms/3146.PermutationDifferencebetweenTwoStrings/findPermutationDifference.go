package findPermutationDifference

import (
	"math"
	"slices"
	"strings"
)

func FindPermutationDifference(s string, t string) int {
	slist := strings.Split(s, "")
	tlist := strings.Split(t, "")
	result := 0
	for i := 0; i < len(tlist); i++ {
		tIndex := slices.Index(slist, tlist[i])
		result += int(math.Abs(float64(i - tIndex)))
	}
	return result
}
