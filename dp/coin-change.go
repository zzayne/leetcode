package dp

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	max := amount + 1

	for j := 1; j <= amount; j++ {
		dp[j] = max
		for i := 0; i < len(coins); i++ {
			if coins[i] <= j && dp[j-coins[i]] != max {
				dp[j] = min(dp[j], dp[j-coins[i]]+1)
			}
		}
	}

	if dp[amount] == max {
		return -1

	}
	return dp[amount]

}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
