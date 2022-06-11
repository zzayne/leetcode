package list

func sortedSquares(nums []int) []int {
	n := len(nums)
	res := make([]int, n)

	l, r := 0, n-1
	for i := n - 1; i >= 0; i-- {
		v1, v2 := nums[l]*nums[l], nums[r]*nums[r]
		if v1 < v2 {
			res[i] = v2
			r--
		} else {
			res[i] = v1
			l++
		}
	}
	return res
}
