package other

func hammingDistance(x, y int) int {
	z := x ^ y

	res := 0
	for z != 0 {
		z &= z - 1
		res++
	}
	return res
}
