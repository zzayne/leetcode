package string

import (
	"math"
)

// https://leetcode.cn/problems/first-unique-character-in-a-string

func firstUniqChar(s string) byte {
	var lst [26]int
	for i, c := range s {
		v := c - 'a'
		if lst[v] == 0 {
			lst[v] = i + 1
		} else {
			lst[v] = -1
		}
	}
	minIdx := math.MaxInt
	var res byte
	for v, idx := range lst {
		if idx == -1 || idx == 0 {
			continue
		}
		if idx-1 < minIdx {
			minIdx = idx
			res = byte(v + 'a')
		}
	}
	if minIdx == math.MaxInt {
		return ' '
	}
	return res
}

func firstUniqChar2(s string) int {
	m := make(map[rune]int)
	for i, v := range []rune(s) {
		if _, ok := m[v]; ok {
			m[v] = -1
			continue
		}
		m[v] = i
	}
	if len(m) == 0 {
		return -1
	}
	min := len(s)
	for _, v := range m {
		if v < min && v != -1 {
			min = v
		}
	}
	if min == len(s) {
		return -1
	}
	return min
}
