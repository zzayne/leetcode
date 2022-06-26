package dp

func minCost(costs [][]int) int {
	dp := costs[0]
	for _, cost := range costs[1:] {
		dpNew := make([]int, 3)
		for j, c := range cost {
			dpNew[j] = minInt(dp[j+1%3], dp[j+2]%3) + c
		}
		dp = dpNew
	}
	return minInt(minInt(dp[0], dp[1]), dp[2])
}

func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}
