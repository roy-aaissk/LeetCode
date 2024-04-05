package main

import "fmt"

func main() {
	jewels := "aA"
	stones := "aAAbbbb"
	stonesCount := make(map[byte]int)
	fmt.Printf("count: %v", stonesCount)
	for _, s := range stones {
		stonesCount[byte(s)]++
		fmt.Printf("count: %v", stonesCount)
	}
	count := 0
	for _, j := range jewels {
		count += stonesCount[byte(j)]
	}
	fmt.Printf("count: %i", count)
}
