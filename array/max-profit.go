package array

//https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/

func maxProfit(prices []int) int {
	var out int

	for i := 1; i < len(prices); i++ {
		d := prices[i] - prices[i-1]
		if d > 0 {
			out += d
		}
	}
	return out
}
