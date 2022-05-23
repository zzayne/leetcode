package string

// https://leetcode.cn/problems/first-unique-character-in-a-string

func firstUniqChar(s string) int {
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
