package string

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}

	l, r, window := 0, 0, make(map[byte]int)
	maxLen := 1

	for r < len(s) {
		rightChar := s[r]
		nextRightIdx := 0

		if val, ok := window[rightChar]; ok {
			nextRightIdx = val
		}
		if nextRightIdx > l {
			l = nextRightIdx
		}
		if r-l+1 > maxLen {
			maxLen = r - l + 1
		}
		window[rightChar] = r + 1
		r++
	}

	return maxLen
}
