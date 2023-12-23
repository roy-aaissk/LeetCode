package Solutions

func theMaximumAchievableX(num int, t int) int {
	achieveNum := num
	for i := 0; i < t+t; i++ {
		achieveNum++
	}
	return achieveNum
}
