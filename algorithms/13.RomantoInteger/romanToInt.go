package romantointeger

func RomanToInt(s string) int {
	if len(s) > 15 || s == "" {
		return 0
	}
	roman_number := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	num, lastint, total := 0, 0, 0
	for i := 0; i < len(s); i++ {
		// 反転して後ろから文字列を判定するとIVを4と評価できる
		char := s[len(s)-(i+1) : len(s)-i]
		num = roman_number[char]
		if num < lastint { // 1 < 5
			total = total - num
		} else { // 5 < 0
			total = total + num
		}
		lastint = num
	}
	return total
}
