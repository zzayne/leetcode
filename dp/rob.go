package dp

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	max := dp[0]
	dp[1] = maxInt(nums[0], nums[1])
	max = maxInt(max, dp[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = maxInt(dp[i-1], dp[i-2]+nums[i])
		max = maxInt(dp[i], max)
	}
	return max

}

func maxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
