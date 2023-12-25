package Solutions

import "strings"

func numJewelsInStones(jewels string, stones string) int {
	stoneSlice := strings.Split(stones, "")
	jewelSlice := strings.Split(jewels, "")
	var result int
	for _, stone := range stoneSlice {
		for i:= 0; i < len(jewelSlice); i++ {
			if jewelSlice[i] == stone {
				result++
			}
		}
	}
	return result
}

// byte型とint型のmapを作りstonesの文字をbyte型をkeyにして分けることでjewelsとstonesを比較するともっと早くなる
