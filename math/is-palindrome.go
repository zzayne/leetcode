package math

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var y int
	tmp := x
	for tmp > 0 {
		y = y*10 + tmp%10
		tmp = tmp / 10
	}

	return y == x
}
