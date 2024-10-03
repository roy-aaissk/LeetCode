package ranktransformofanarray

import "sort"

func ArrayRankTransform(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}

	// Step 1: Create a sorted copy of the array
	sortedArr := make([]int, len(arr))
	copy(sortedArr, arr)
	// sort.Intsを使う方が高速
	sort.Ints(sortedArr)

	// Step 2: Create a map to assign ranks
	rankMap := make(map[int]int)
	rank := 1

	// Step 3: Assign ranks to sorted elements
	for _, num := range sortedArr {
		if _, exists := rankMap[num]; !exists {
			rankMap[num] = rank
			rank++
		}
	}

	// Step 4: Replace each element in the original array with its rank
	for i, num := range arr {
		arr[i] = rankMap[num]
	}

	return arr
}

// bubbleSortに時間がかかる
// func bubbleSort(arr []int) {
// 	n := len(arr)
// 	for i := 0; i < n-1; i++ {
// 		for j := 0; j < n-i-1; j++ {
// 			if arr[j] > arr[j+1] {
// 				arr[j], arr[j+1] = arr[j+1], arr[j]
// 			}
// 		}
// 	}
// }

// func ArrayRankTransform(arr []int) []int {
// 	rankResult := []int{}
// 	sortArr := make([]int, len(arr))
// 	copy(sortArr, arr)
// 	// 並べ替え
// 	bubbleSort(sortArr)
// 	// rankづけのmapを作成
// 	// 重複の登録はしない
// 	rankMap := make(map[int]int)
// 	rank := 1
// 	for _, num := range sortArr {
// 		if _, exist := rankMap[num]; !exist {
// 			rankMap[num] = rank
// 			rank++
// 		}
// 	}
// 	for _, n := range arr {
// 		rankResult = append(rankResult, rankMap[n])
// 	}
// 	return rankResult
// }
