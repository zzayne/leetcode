package array

// https://leetcode.cn/problems/plus-one/

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i > -1; i-- {
		digits[i]++
		if digits[i] != 10 {
			return digits
		}
		digits[i] = 0
	}
	return append([]int{1}, digits...)

}
