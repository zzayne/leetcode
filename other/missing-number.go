package other

func missingNumber(nums []int) int {
	length := len(nums)
	sum := (0 + length) * (length + 1) / 2
	for _, v := range nums {
		sum -= v
	}
	return sum
}

func missingNumber2(nums []int) int {
	var res int
	for i, v := range nums {
		res ^= v ^ (i + 1)
	}
	return res
}
