package dp

func maxSubArray(nums []int) int {
	cur := nums[0]
	max := cur
	for i, v := range nums {
		if i == 0 {
			continue
		}
		if cur < 0 {
			cur = v
		} else {
			cur = cur + v
		}
		if cur > max {
			max = cur
		}
	}
	return max
}
