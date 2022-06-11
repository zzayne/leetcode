package list

func rotate(nums []int, k int) {
	n := len(nums)

	if k == 0 || k%n == 0 {
		return
	}
	newNums := make([]int, n)
	for i, v := range nums {
		newNums[(i+k)%n] = v
	}
	copy(nums, newNums)
	return
}
