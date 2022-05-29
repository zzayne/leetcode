package string

import (
	"math"
)

//https://leetcode.cn/problems/string-to-integer-atoi

func myAtoi(s string) int {
	var res int32
	sign := 1
	idx := 0
	lst := []rune(s)
	for idx < len(lst) && lst[idx] == 32 {
		idx++
	}

	if idx == len(lst) {
		return 0
	}

	if lst[idx] == 43 || lst[idx] == 45 {
		if lst[idx] == 45 {
			sign = -1
		}
		idx++
	}

	for idx < len(lst) {
		v := lst[idx]
		if v < 48 || v > 57 {
			break
		}
		temp := res
		res = res*10 + (int32(v) - 48)
		if res/10 != temp {
			if sign > 0 {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}
		idx++
	}
	return sign * int(res)
}
