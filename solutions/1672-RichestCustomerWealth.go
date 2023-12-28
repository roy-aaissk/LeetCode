package Solutions

func maximumWealth(accounts [][]int) int {
	max := 0
	var addAccounts []int
	for _, account := range accounts {
		var total int
		for i := 0; i < len(account); i++ {
			total = total + account[i]
		}
		addAccounts = append(addAccounts, total)
	}
	for _, addAccount := range addAccounts {
		if max < addAccount {
			max = addAccount
		}
	}
	return max
}
