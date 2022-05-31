package math

//https://leetcode.cn/problems/power-of-three/

func isPowerOfThree(n int) bool {
	return n > 0 && (n == 1 || (n%3 == 0 && isPowerOfThree(n/3)))
}
