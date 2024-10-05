package main

import (
	"fmt"
	romantointeger "leet_code/algorithms/13.RomantoInteger"
	ranktransformofanarray "leet_code/algorithms/1331.RankTransformofanArray"
	concatenationofarray "leet_code/algorithms/1929.ConcatenationofArray"
	scoreofastring "leet_code/algorithms/3110.ScoreofaString"
)

func main() {
	fmt.Print(romantointeger.RomanToInt("MMMCCCXXXIV"))
	fmt.Print(ranktransformofanarray.ArrayRankTransform([]int{40, 10, 20, 30}))
	fmt.Printf("\n result: %v \n", concatenationofarray.GetConcatenation([]int{1, 2, 1}))
	fmt.Printf("\n scoreOfString:%v\n", scoreofastring.ScoreOfString("hello"))
}
