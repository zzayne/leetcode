package array

// https://leetcode.cn/problems/move-zeroes

func moveZeroes(nums []int) {
	var j int
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			j++
		} else if j != 0 {
			nums[i-j] = nums[i]
			nums[i] = 0
		}
	}
	return
}
