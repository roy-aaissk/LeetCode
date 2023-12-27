package Solutions

import "strings"

func findWordsContaining(words []string, x byte) []int {
	var result []int
	for i, word := range words {
		if strings.Contains(word, string(x)) {
			result = append(result, i)
			continue
		}
	}
	return result
}
