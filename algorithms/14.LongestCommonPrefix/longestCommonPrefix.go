package longestcommonprefix

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	// Start with the first string as the prefix
	prefix := strs[0]

	// Compare the prefix with each string in the array
	for i := 1; i < len(strs); i++ {
		// Update the prefix by comparing it with the current string
		prefix = commonPrefix(prefix, strs[i])
		// If at any point the prefix is empty, there is no common prefix
		if prefix == "" {
			break
		}
	}
	return prefix
}

// Helper function to find the common prefix between two strings
func commonPrefix(str1, str2 string) string {
	minLength := min(len(str1), len(str2))
	for i := 0; i < minLength; i++ {
		if str1[i] != str2[i] {
			return str1[:i]
		}
	}
	return str1[:minLength]
}

// Helper function to find the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
