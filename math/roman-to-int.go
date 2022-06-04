package math

func romanToInt(s string) int {
	m := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	var res int
	lst := []rune(s)
	for i, v := range lst {
		if i == len(s)-1 {
			res += m[v]
			return res
		}
		next := lst[i+1]
		if v == 'I' && (next == 'V' || next == 'X') {
			res -= m[v]
			continue
		}
		if v == 'X' && (next == 'L' || next == 'C') {
			res -= m[v]
			continue
		}
		if v == 'C' && (next == 'D' || next == 'M') {
			res -= m[v]
			continue
		}
		res += m[v]
	}
	return res
}
