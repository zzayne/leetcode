package other

func hammingWeight(num uint32) int {
	i := 0
	for num != 0 {
		if num%2 == 1 {
			i++
		}
		num = num >> 1
	}
	return i
}
