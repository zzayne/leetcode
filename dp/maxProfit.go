package dp

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	hold := -prices[0]
	noHold := 0
	for i := 1; i < len(prices); i++ {
		if hold+prices[i] > noHold {
			noHold = hold + prices[i]
		}
		if -prices[i] > hold {
			hold = -prices[i]
		}
	}
	return noHold
}

func maxProfit2(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	max := 0
	min := prices[0]
	for i := 0; i < len(prices); i++ {
		if min > prices[i] {
			min = prices[i]
		}
		if prices[i]-min > max {
			max = prices[i] - min
		}
	}
	return max
}
