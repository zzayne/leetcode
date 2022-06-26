package math

func fib(n int) int {
	a, b := 0, 1
	if n == 0 {
		return a
	}
	if n == 1 {
		return b
	}
	var tmp int
	for i := 2; i <= n; i++ {
		tmp = b
		b = a + b
		a = tmp
	}
	return b
}
