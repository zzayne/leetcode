package math

//https://leetcode.cn/problems/count-primes/

// Sieve_of_Eratosthenes https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes
// Time Complexity：O(nlogn)
// Space Complexity：O（n）
func countPrimes(n int) int {
	arr := make([]bool, n)
	res := 0

	for i := 2; i < n; i++ {
		if arr[i] {
			continue
		}
		res++
		for j := i; j < n; j += i {
			arr[j] = true
		}
	}
	return res
}
