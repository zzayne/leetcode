package array

import "math"

func find132pattern(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}

	var stack []int
	k := math.MinInt

	for i := n - 1; i > 0; i-- {
		if nums[i] < k {
			return true
		}

		for len(stack) != 0 && stack[len(stack)-1] < nums[i] {
			k = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums[i])
	}
	return false

}
